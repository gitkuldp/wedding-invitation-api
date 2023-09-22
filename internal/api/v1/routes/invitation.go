package v1

import (
	v1 "github.com/gitkuldp/wedding-invitation-api/internal/api/v1/handlers"
	"github.com/gitkuldp/wedding-invitation-api/internal/services"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterInvitationRoute(db *gorm.DB, group *echo.Group) {
	invitationService := services.NewInvitationService(db)
	ih := v1.NewInvitationHandler(invitationService)
	group.POST("/invitations", ih.Create)
	group.GET("/invitations", ih.List)
	group.POST("/invitations/:id", ih.Update)
	group.GET("/invitations/:id", ih.Find)
}
