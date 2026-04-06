package enclave

import (
	"context"
	"errors"
	"reflect"
	"testing"
)

func TestIteratorCollectAllPages(t *testing.T) {
	ctx := context.Background()
	calls := 0

	it := newIterator(ctx, 2, func(_ context.Context, limit, offset int) ([]int, error) {
		calls++
		if limit != 2 {
			t.Fatalf("unexpected limit: got %d", limit)
		}
		switch offset {
		case 0:
			return []int{1, 2}, nil
		case 2:
			return []int{3}, nil
		default:
			return nil, nil
		}
	})

	var got []int
	for it.Next() {
		got = append(got, it.Value())
	}
	if err := it.Err(); err != nil {
		t.Fatalf("unexpected iterator error: %v", err)
	}
	if !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Fatalf("unexpected items: got %v", got)
	}
	if calls != 2 {
		t.Fatalf("unexpected fetch count: got %d want 2", calls)
	}
}

func TestIteratorPropagatesFetcherError(t *testing.T) {
	expectedErr := errors.New("boom")
	it := newIterator(context.Background(), 10, func(_ context.Context, _, _ int) ([]int, error) {
		return nil, expectedErr
	})

	if it.Next() {
		t.Fatalf("expected Next to be false on fetch error")
	}
	if !errors.Is(it.Err(), expectedErr) {
		t.Fatalf("expected error %v, got %v", expectedErr, it.Err())
	}
}

func TestCollectAll(t *testing.T) {
	it := newIterator(context.Background(), 3, func(_ context.Context, _, offset int) ([]int, error) {
		if offset == 0 {
			return []int{10, 11}, nil
		}
		return nil, nil
	})

	all, err := collectAll(it)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(all, []int{10, 11}) {
		t.Fatalf("unexpected result: got %v", all)
	}
}

func TestEnumerateSendsError(t *testing.T) {
	expectedErr := errors.New("fetch failed")
	it := newIterator(context.Background(), 2, func(_ context.Context, _, _ int) ([]int, error) {
		return nil, expectedErr
	})

	ch := enumerate(context.Background(), it)
	res, ok := <-ch
	if !ok {
		t.Fatalf("expected one result item")
	}
	if !errors.Is(res.Err, expectedErr) {
		t.Fatalf("expected %v, got %v", expectedErr, res.Err)
	}
	if _, ok := <-ch; ok {
		t.Fatalf("expected channel to be closed")
	}
}
