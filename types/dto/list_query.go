package dto

type ListQuery struct {
	Filter string
}

type PaginatedListQuery struct {
	ListQuery
	PaginationQuery
}
