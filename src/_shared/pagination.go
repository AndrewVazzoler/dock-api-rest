package shared

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type QueryPagination struct {
	Limit int    `form:"limit"`
	Page  int    `form:"page"`
	Sort  string `form:"sort"`
}

type Pagination[T any] struct {
	Limit      int    `json:"limit,omitempty;query:limit"`
	Page       int    `json:"page,omitempty;query:page"`
	Sort       string `json:"sort,omitempty;query:sort"`
	TotalRows  int64  `json:"total_rows"`
	TotalPages int    `json:"total_pages"`
	Rows       T      `json:"rows"`
}

func (p *Pagination[any]) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination[any]) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination[any]) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination[any]) GetSort() string {
	if p.Sort == "" {
		p.Sort = "created_at asc"
	}
	return p.Sort
}

func Paginate[T any](value interface{}, pagination *Pagination[T], db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)

	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit())))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}

func GetQueryPaginate(c *gin.Context) QueryPagination {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "0"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))
	return QueryPagination{
		Limit: limit,
		Page:  page,
		Sort:  c.DefaultQuery("sort", ""),
	}
}
