package enclave

import "context"

type pageFetcher[T any] func(ctx context.Context, limit int, offset int) ([]T, error)

type Iterator[T any] struct {
	ctx      context.Context
	pageSize int
	fetch    pageFetcher[T]

	buffer []T
	index  int
	offset int
	value  T
	err    error
	done   bool
	final  bool
}

type ItemResult[T any] struct {
	Item T
	Err  error
}

func newIterator[T any](ctx context.Context, pageSize int, fetch pageFetcher[T]) *Iterator[T] {
	return &Iterator[T]{
		ctx:      ctx,
		pageSize: pageSize,
		fetch:    fetch,
	}
}

func (it *Iterator[T]) Next() bool {
	if it.done {
		return false
	}
	if it.err != nil {
		it.done = true
		return false
	}
	if err := it.ctx.Err(); err != nil {
		it.err = err
		it.done = true
		return false
	}

	if it.index >= len(it.buffer) {
		items, err := it.fetch(it.ctx, it.pageSize, it.offset)
		if err != nil {
			it.err = err
			it.done = true
			return false
		}
		if len(items) == 0 {
			it.done = true
			return false
		}
		it.buffer = items
		it.index = 0
		it.offset += len(items)
		if len(items) < it.pageSize {
			// Mark that no additional page fetch is needed after current buffer is drained.
			it.final = true
		}
	}

	it.value = it.buffer[it.index]
	it.index++

	if it.index < len(it.buffer) {
		return true
	}
	if it.final {
		// Buffer drained and no additional pages are expected.
		it.done = true
		it.buffer = nil
		it.index = 0
	}
	return true
}

func (it *Iterator[T]) Value() T {
	return it.value
}

func (it *Iterator[T]) Err() error {
	return it.err
}

func enumerate[T any](ctx context.Context, it *Iterator[T]) <-chan ItemResult[T] {
	ch := make(chan ItemResult[T])
	go func() {
		defer close(ch)
		for it.Next() {
			select {
			case <-ctx.Done():
				return
			case ch <- ItemResult[T]{Item: it.Value()}:
			}
		}
		if err := it.Err(); err != nil {
			select {
			case <-ctx.Done():
			case ch <- ItemResult[T]{Err: err}:
			}
		}
	}()
	return ch
}

func collectAll[T any](it *Iterator[T]) ([]T, error) {
	items := make([]T, 0)
	for it.Next() {
		items = append(items, it.Value())
	}
	if err := it.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
