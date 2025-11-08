package pagination

import (
	"errors"
	"math"
)

var (
	ErrPageOutOfRange = errors.New("page out of range")
	ErrSizeOutOfRange = errors.New("size out of range")
)

type Pagination struct {
	Page    int    `json:"page"`
	Size    int    `json:"size"`
	OrderBy string `json:"order_by"`
	SortBy  string `json:"sort_by"`
}

type Query struct {
	Offset     int
	Limit      int
	TotalPages int
	OrderBy    string
	SortBy     string
	NextPage   int
	PrevPage   int
}

type Response[T any] struct {
	Pagination struct {
		Page       int    `json:"page"`
		Size       int    `json:"size"`
		TotalPages int    `json:"total_pages"`
		TotalItems int    `json:"total_items"`
		NextPage   int    `json:"next_page"`
		PrevPage   int    `json:"prev_page"`
		OrderBy    string `json:"order_by"`
		SortBy     string `json:"sort_by"`
	} `json:"pagination"`
	Data T `json:"data"`
}

func (p *Pagination) Values(totalItems int) (Query, error) {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Size <= 0 {
		return Query{}, ErrSizeOutOfRange
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(p.Size)))
	if totalPages == 0 {
		totalPages = 1
	}
	if totalItems > 0 && p.Page > totalPages {
		return Query{}, ErrPageOutOfRange
	}

	orderBy := defaultStr(p.OrderBy, "created_at")
	sortBy := defaultStr(p.SortBy, "DESC")

	return Query{
		Offset:     (p.Page - 1) * p.Size,
		Limit:      p.Size,
		TotalPages: totalPages,
		OrderBy:    orderBy,
		SortBy:     sortBy,
		NextPage:   min(p.Page+1, totalPages),
		PrevPage:   max(p.Page-1, 1),
	}, nil
}

func defaultStr(s, def string) string {
	if s == "" {
		return def
	}
	return s
}

func min(a, b int) int { if a < b { return a }; return b }
func max(a, b int) int { if a > b { return a }; return b }
