package modules

import (
	"rest-api/api"
	"rest-api/api/v1/auth"
	"rest-api/api/v1/movie"
	"rest-api/api/v1/user"
	movieService "rest-api/business/movie"
	authService "rest-api/business/user"
	jwtService "rest-api/business/user"
	userService "rest-api/business/user"
	"rest-api/config"
	movieRepo "rest-api/repository/movie"
	userRepo "rest-api/repository/user"
	"rest-api/util"
)

func RegisterModules(dbCon *util.DatabaseConnection, config *config.AppConfig) api.Controller {
	userRepo := userRepo.RepositoryFactory(dbCon)

	userService := userService.NewUserService(userRepo)
	authService := authService.NewAuthService(userRepo)
	jwtService := jwtService.NewJWTService()

	movieRepo := movieRepo.MovieRepositoryFactory(dbCon)
	movieService := movieService.NewMovieService(movieRepo)

	controller := api.Controller{
		Auth:  auth.NewAuthController(authService, jwtService, userService),
		User:  user.NewUserController(userService, jwtService),
		Movie: movie.NewMovieController(movieService, jwtService),
	}

	return controller
}
