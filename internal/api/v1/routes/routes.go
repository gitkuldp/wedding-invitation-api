package v1

import (
	"github.com/gitkuldp/wedding-invitation-api/internal/db"

	"github.com/labstack/echo/v4"
)

func Setup(app *db.Application, echo *echo.Echo) {
	erp := echo.Group("api/v1/erp")
	// authGroup := echo.Group("api/v1/erp")

	// authGroup.Use(middleware.Authentication(app.Db))
	RegisterUserRoute(app.Db, erp)
	LoginUserRoute(app.Db, erp)
	RegisterInvitationRoute(app.Db, erp)
	BrideRoute(app.Db, erp)
	GroomRoute(app.Db, erp)
	RegisterAttachmentRoute(app.Db, erp)
}
