package v1

import (
	"github.com/gitkuldp/wedding-invitation-api/internal/api/middleware"
	v1 "github.com/gitkuldp/wedding-invitation-api/internal/api/v1/handlers"
	"github.com/gitkuldp/wedding-invitation-api/internal/services"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterAttachmentRoute(db *gorm.DB, group *echo.Group) {
	attachment := services.NewAttachmentService(db)
	ah := v1.NewAttachmentHandler(attachment)

	group.POST("/attachments", ah.Create)
	group.GET("/attachments", ah.List, middleware.GetParams)
	group.GET("/attachments/:id", ah.Detail)
}
