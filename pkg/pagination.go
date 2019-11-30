package pkg

// Pagination 提供给前端的页码生成器
type Pagination struct {
	Total    uint `json:"total"`
	PageSize uint `json:"pageSize"`
	Current  uint `json:"current"`
	Offset   uint `json:"-"`
}
