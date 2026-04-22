package enclave

import "iter"

// defaultPageSize is the number of items fetched per
// page in paginated list requests.
const defaultPageSize = 50

// Collect drains an iter.Seq2[T, error] into a slice,
// stopping at the first error.
func Collect[T any](
	seq iter.Seq2[T, error],
) ([]T, error) {
	var result []T

	for item, err := range seq {
		if err != nil {
			return result, err
		}

		result = append(result, item)
	}

	return result, nil
}

// paginate returns an iter.Seq2 that transparently
// fetches pages using the provided fetchPage function.
// fetchPage receives limit and offset and returns a
// page of items. Iteration stops when a page returns
// fewer items than the limit.
func paginate[T any](
	fetchPage func(
		limit int,
		offset int,
	) ([]T, error),
) iter.Seq2[T, error] {
	return func(yield func(T, error) bool) {
		offset := 0

		for {
			page, err := fetchPage(
				defaultPageSize,
				offset,
			)
			if err != nil {
				var zero T
				yield(zero, err)

				return
			}

			for _, item := range page {
				if !yield(item, nil) {
					return
				}
			}

			if len(page) < defaultPageSize {
				return
			}

			offset += len(page)
		}
	}
}
