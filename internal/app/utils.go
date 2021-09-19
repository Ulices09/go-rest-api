package app

import (
	"go-rest-api/internal/types/dto"
	"go-rest-api/internal/types/entity"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetLoggedInUser(c echo.Context) entity.Claims {
	claims := c.Get("user").(*entity.Claims)
	return *claims
}

func GetListQuery(c echo.Context) (query dto.ListQuery) {
	filter := c.QueryParam("filter")

	query = dto.ListQuery{
		Filter: filter,
	}

	return
}

func GetListPaginatedQuery(c echo.Context) (query dto.PaginatedListQuery, err error) {
	paginationQuery, err := GetPaginationQuery(c)

	if err != nil {
		return
	}

	filter := c.QueryParam("filter")

	query = dto.PaginatedListQuery{
		ListQuery: dto.ListQuery{
			Filter: filter,
		},
		PaginationQuery: paginationQuery,
	}

	return
}

func GetPaginationQuery(c echo.Context) (query dto.PaginationQuery, err error) {
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

	query = dto.PaginationQuery{
		Page: page,
		Skip: skip,
		Take: take,
	}

	return
}
