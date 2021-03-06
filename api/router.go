package api

import (
	"rest-api/api/middleware"
	"rest-api/api/v1/auth"
	"rest-api/api/v1/movie"
	"rest-api/api/v1/user"

	service "rest-api/business/user"

	"github.com/labstack/echo/v4"
)

var jwtService service.JWTService = service.NewJWTService()

type Controller struct {
	Auth  *auth.AuthController
	User  *user.UserController
	Movie *movie.MovieController
}

func RegisterRoutes(e *echo.Echo, controller *Controller) {

	authRoutes := e.Group("/api/v1/auth")
	authRoutes.POST("/login", controller.Auth.Login)
	authRoutes.POST("/register", controller.Auth.RegisterHandler)

	userRoutes := e.Group("/api/v1/user", middleware.AuthorizeJWT(jwtService))
	userRoutes.GET("/profile", controller.User.Profile)
	userRoutes.PUT("/profile", controller.User.Update)

	movieRoutes := e.Group("/api/v1/movie", middleware.AuthorizeJWT(jwtService))
	movieRoutes.GET("/mylist", controller.Movie.All)
	movieRoutes.GET("", controller.Movie.SearchMovie)
	movieRoutes.POST("/addlist", controller.Movie.AddWishList)
	movieRoutes.GET("/mylist/:id", controller.Movie.FindOneMovieByID)
	movieRoutes.DELETE("/mylist/:id", controller.Movie.DeleteMovie)
}
