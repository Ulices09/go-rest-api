package app

import (
	"go-rest-api/types/dto"
	"go-rest-api/types/entity"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetLoggedInUser(c echo.Context) entity.Claims {
	claims := c.Get("user").(*entity.Claims)
	return *claims
}

func GetPaginationQuery(c echo.Context) (query *dto.PaginationQuery, err error) {
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
			return nil, err
		}

		if page > 0 {
			if pageSizeQuery != "" {
				take, err = strconv.Atoi(pageSizeQuery)

				if err != nil {
					return nil, err
				}
			}

			skip = take * (page - 1)
		}
	}

	query = &dto.PaginationQuery{
		Page: page,
		Skip: skip,
		Take: take,
	}

	return query, nil
}
