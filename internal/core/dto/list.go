package dto

import (
	"strconv"

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

func NewPaginationQuery(c echo.Context) (query PaginationQuery, err error) {
	pageQuery := c.QueryParam("page")
	pageSizeQuery := c.QueryParam("pageSize")

	var (
		page = 0
		skip = 0
		take = 10
	)

	if pageQuery != "" {
		page, err = strconv.Atoi(pageQuery)

		if err != nil {
			return
		}

		if page > 0 {
			if pageSizeQuery != "" {
				take, err = strconv.Atoi(pageSizeQuery)

				if err != nil {
					return
				}
			}

			skip = take * (page - 1)
		}
	}

	query = PaginationQuery{
		Page: page,
		Skip: skip,
		Take: take,
	}

	return
}

type ListResult struct {
	Data interface{} `json:"data"`
}
