package services

import (
	"math"

	"github.com/gitkuldp/wedding-invitation-api/internal/models"
	"github.com/gitkuldp/wedding-invitation-api/internal/utils"
	"gorm.io/gorm"
)

func ApplyQueryParams(db *gorm.DB, params models.QueryParams) *gorm.DB {
	if params.Pagination {
		db = db.Offset((params.Page - 1) * params.Limit)
		db = db.Limit(params.Limit)
	}
	// FIXME: inactive is not a common column
	if params.Inactive != nil {
		db = db.Where("is_active = ?", params.Inactive)
	}

	if params.TableName != "" {
		db = db.Where("table_name = ?", params.TableName)
	}
	if params.TableId != "" {
		db = db.Where("table_id = ?", params.TableId)
	}

	return db
}

func CalculateMetaData(db *gorm.DB, params models.QueryParams, model interface{}) *models.Pagination {
	var totalCount int64
	//var currentCount int64
	var pagination models.Pagination

	//db.Model(model).Count(&currentCount)
	db.Model(model).Offset(-1).Limit(-1).Count(&totalCount)

	pagination.CurrentPage = params.Page
	pagination.Limit = params.Limit
	pagination.TotalCount = totalCount
	pagination.TotalPage = int(math.Ceil(float64(totalCount) / float64(params.Limit)))
	//pagination.CurrentCount = currentCount

	return &pagination
}

func Paginate(db *gorm.DB, pagination utils.Pagination, currentCount int, model interface{}) (*models.Pagination, error) {
	if !pagination.Paginate {
		return nil, nil
	}

	var count int64
	err := db.Model(model).Offset(-1).Limit(-1).Count(&count).Error
	if err != nil {
		return nil, err
	}

	return &models.Pagination{
		CurrentPage:  pagination.Page,
		TotalPage:    pagination.TotalPage(count),
		Limit:        pagination.Limit,
		CurrentCount: currentCount,
		TotalCount:   count,
	}, nil
}
