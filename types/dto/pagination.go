package dto

import "math"

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
