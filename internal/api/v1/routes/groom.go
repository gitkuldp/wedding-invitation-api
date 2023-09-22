package v1

import (
	v1 "github.com/gitkuldp/wedding-invitation-api/internal/api/v1/handlers"
	"github.com/gitkuldp/wedding-invitation-api/internal/services"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GroomRoute(db *gorm.DB, group *echo.Group) {
	groomService := services.NewGroomService(db)
	ih := v1.NewGroomHandler(groomService)
	group.POST("/grooms", ih.Create)
	group.GET("/grooms", ih.List)
	group.POST("/grooms/:id", ih.Update)
	group.GET("/grooms/:id", ih.Find)
}
