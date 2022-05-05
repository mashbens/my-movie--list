package api

import (
	"rest-api/api/middleware"
	"rest-api/api/v1/auth"
	"rest-api/api/v1/user"

	service "rest-api/business/user"

	"github.com/labstack/echo/v4"
)

var jwtService service.JWTService = service.NewJWTService()

type Controller struct {
	Auth *auth.AuthController
	User *user.UserController
}

func RegisterRoutes(e *echo.Echo, controller *Controller) {

	authRoutes := e.Group("/api/v1/auth")
	authRoutes.POST("/login", controller.Auth.Login)
	authRoutes.POST("/register", controller.Auth.Register)

	userRoutes := e.Group("/api/v1/user", middleware.AuthorizeJWT(jwtService))
	userRoutes.GET("/profile", controller.User.Profile)
	userRoutes.PUT("/profile", controller.User.Update)

}
