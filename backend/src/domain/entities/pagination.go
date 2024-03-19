package entities

type Pagination[T any] struct {
	Page        int
	RowsPerPage int
	TotalRows   int
	Results     []*T
}
