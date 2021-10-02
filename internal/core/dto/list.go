package dto

import (
	"github.com/labstack/echo/v4"
)

type ListQuery struct {
	Filter string
}

func NewListQuery(c echo.Context) (query ListQuery) {
	filter := c.QueryParam("filter")

	query = ListQuery{
		Filter: filter,
	}

	return
}

type PaginatedListQuery struct {
	ListQuery
	PaginationQuery
}

func NewListPaginatedQuery(c echo.Context) (query PaginatedListQuery, err error) {
	paginationQuery, err := NewPaginationQuery(c)

	if err != nil {
		return
	}

	filter := c.QueryParam("filter")

	query = PaginatedListQuery{
		ListQuery: ListQuery{
			Filter: filter,
		},
		PaginationQuery: paginationQuery,
	}

	return
}

type ListResult struct {
	Data interface{} `json:"data"`
}
