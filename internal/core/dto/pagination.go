package dto

import (
	"math"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PaginationQuery struct {
	Page int
	Skip int
	Take int
}

type PaginationResult struct {
	Data     interface{} `json:"data"`
	Total    int         `json:"total"`
	Pages    int         `json:"pages"`
	PageSize int         `json:"pageSize"`
}

func NewPaginationResult(data interface{}, total int, take int) PaginationResult {
	pages := 0

	if take > 0 {
		div := float64(total) / float64(take)
		pages = int(math.Ceil(div))
	}

	result := PaginationResult{
		Data:     data,
		Total:    total,
		Pages:    pages,
		PageSize: take,
	}

	return result
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
