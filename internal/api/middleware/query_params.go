package middleware

import (
	"strconv"

	"github.com/gitkuldp/wedding-invitation-api/internal/models"
	"github.com/gitkuldp/wedding-invitation-api/internal/utils"
	"github.com/labstack/echo/v4"
)

func GetParams(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var params models.QueryParams

		// Parse "pagination" parameter
		params.Pagination = true
		if paginationParam := c.QueryParam("pagination"); paginationParam == "false" {
			params.Pagination = false
		}

		// Parse "page" parameter
		params.Page = 1
		if pageParam := c.QueryParam("page"); pageParam != "" {
			if page, err := strconv.Atoi(pageParam); err == nil && page >= 1 {
				params.Page = page
			}
		}

		// Parse "limit" parameter
		params.Limit = 10
		if limitParam := c.QueryParam("limit"); limitParam != "" {
			if limit, err := strconv.Atoi(limitParam); err == nil && limit >= 10 && limit <= 1000 {
				params.Limit = limit
			}
		}

		// Parse "inactive" parameter
		inactiveParam := c.QueryParam("is_active")
		if inactiveParam == "true" {
			params.Inactive = utils.Ref(true)
		} else if inactiveParam == "false" {
			params.Inactive = utils.Ref(false)
		}

		// Parse "status" parameter
		params.TableId = c.QueryParam("table_id")

		// Parse "status_for" parameter
		params.TableName = c.QueryParam("table_name")

		c.Set("query-params", params)
		return next(c)
	}
}
