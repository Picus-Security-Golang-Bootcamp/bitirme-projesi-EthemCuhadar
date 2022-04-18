package helper

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// Pagination struct to save query parameters from URL.
// Example: "api/v1/product?limit=10&page=2&order=price asc"
type Pagination struct {
	Limit int64  `json:"limit"`
	Page  int64  `json:"page"`
	Sort  string `json:"sort"`
}

// GeneratePaginationFromRequest takes query parameters from URL request
// and save them into Pagination struct. Afterwards, it returns Pagination struct.
func GeneratePaginationFromRequest(c *gin.Context) *Pagination {

	// Default parameters
	limit := int64(10)
	page := int64(1)
	sort := "created_at asc"

	// Query parameters
	limit, _ = strconv.ParseInt(c.Query("limit"), 10, 64)
	page, _ = strconv.ParseInt(c.Query("page"), 10, 64)
	sort = c.Query("order")

	return &Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}
}
