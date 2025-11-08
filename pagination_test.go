package pagination

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaginationValues(t *testing.T) {
	tests := []struct {
		name       string
		pagination Pagination
		totalItems int
		expected   Query
		expectErr  error
	}{
		{
			"Valid pagination",
			Pagination{Page: 1, Size: 10, OrderBy: "description", SortBy: "ASC"},
			100,
			Query{Offset: 0, Limit: 10, TotalPages: 10, OrderBy: "description", SortBy: "ASC", NextPage: 2, PrevPage: 1},
			nil,
		},
		{
			"Default order and sort",
			Pagination{Page: 1, Size: 10},
			100,
			Query{Offset: 0, Limit: 10, TotalPages: 10, OrderBy: "created_at", SortBy: "DESC", NextPage: 2, PrevPage: 1},
			nil,
		},
		{
			"Zero totalItems",
			Pagination{Page: 1, Size: 10},
			0,
			Query{Offset: 0, Limit: 10, TotalPages: 1, OrderBy: "created_at", SortBy: "DESC", NextPage: 1, PrevPage: 1},
			nil,
		},
		{
			"Page out of range",
			Pagination{Page: 5, Size: 10},
			20,
			Query{},
			ErrPageOutOfRange,
		},
		{
			"Zero page",
			Pagination{Page: 0, Size: 10},
			100,
			Query{Offset: 0, Limit: 10, TotalPages: 10, OrderBy: "created_at", SortBy: "DESC", NextPage: 2, PrevPage: 1},
			nil,
		},
		{
			"Negative page",
			Pagination{Page: -1, Size: 10},
			100,
			Query{Offset: 0, Limit: 10, TotalPages: 10, OrderBy: "created_at", SortBy: "DESC", NextPage: 2, PrevPage: 1},
			nil,
		},
		{
			"Zero size",
			Pagination{Page: 1, Size: 0},
			100,
			Query{},
			ErrSizeOutOfRange,
		},
		{
			"Negative size",
			Pagination{Page: 1, Size: -10},
			100,
			Query{},
			ErrSizeOutOfRange,
		},
		{
			"Next and prev pages",
			Pagination{Page: 2, Size: 10},
			100,
			Query{Offset: 10, Limit: 10, TotalPages: 10, OrderBy: "created_at", SortBy: "DESC", NextPage: 3, PrevPage: 1},
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.pagination.Values(tt.totalItems)
			assert.Equal(t, tt.expectErr, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}
