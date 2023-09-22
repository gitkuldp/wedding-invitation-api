package v1

import (
	v1 "github.com/gitkuldp/wedding-invitation-api/internal/api/v1/handlers"
	"github.com/gitkuldp/wedding-invitation-api/internal/services"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterUserRoute(db *gorm.DB, group *echo.Group) {
	userService := services.NewUserService(db)
	rh := v1.NewRegisterHandler(userService)
	group.GET("/users", rh.GetUsers)
}

func LoginUserRoute(db *gorm.DB, group *echo.Group) {
	userService := services.NewUserService(db)
	rh := v1.NewRegisterHandler(userService)
	group.POST("/refresh-tokens", rh.RefreshToken)
	group.POST("/user-registration", rh.RegisterUser)
	group.POST("/login", rh.UserLogin)
}
