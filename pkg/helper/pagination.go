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
	limit := int64(2)
	page := int64(1)
	sort := "created_at asc"

	// Query
	query := c.Request.URL.Query()

	// Assign parameters
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.ParseInt(queryValue, 10, 64)
			break
		case "page":
			page, _ = strconv.ParseInt(queryValue, 10, 64)
			break
		case "sort":
			sort = queryValue
			break
		}
	}

	return &Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}
}
