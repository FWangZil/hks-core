package safeservice

import (
	"hks/hks-core/pkg"
	"hks/hks-core/util"

	"github.com/gin-gonic/gin"
)

// Page PageSize 每页数据条数
const (
	Page     = 1
	PageSize = 12
)

// paginationByGet 通过当前页码新建一个Pagination
func paginationByGet(c *gin.Context) *pkg.Pagination {
	page := c.Query("page")
	pageSize := c.Query("pageSize")
	var p uint = Page
	if page != "" {
		temp, err := util.ParseUint(page)
		if err == nil && temp >= 1 {
			p = temp
		}
	}
	var ps uint = PageSize
	if pageSize != "" {
		temp, err := util.ParseUint(pageSize)
		if err == nil && temp >= 1 {
			ps = temp
		}
	}
	return &pkg.Pagination{
		Total:    0,
		PageSize: ps,
		Current:  p,
		Offset:   (p - 1) * ps,
	}
}
