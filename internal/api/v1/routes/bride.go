package v1

import (
	v1 "github.com/gitkuldp/wedding-invitation-api/internal/api/v1/handlers"
	"github.com/gitkuldp/wedding-invitation-api/internal/services"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func BrideRoute(db *gorm.DB, group *echo.Group) {
	brideService := services.NewBrideService(db)
	ih := v1.NewBrideHandler(brideService)
	group.POST("/bride", ih.Create)
	group.GET("/bride", ih.List)
	group.POST("/bride/:id", ih.Update)
	group.GET("/bride/:id", ih.Find)
}
