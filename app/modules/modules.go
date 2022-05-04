package modules

import (
	"rest-api/api"
	"rest-api/api/v1/auth"
	"rest-api/api/v1/user"
	authService "rest-api/business/user"
	jwtService "rest-api/business/user"
	userService "rest-api/business/user"
	"rest-api/config"
	userRepo "rest-api/repository/user"
	"rest-api/util"
)

func RegisterModules(dbCon *util.DatabaseConnection, config *config.AppConfig) api.Controller {
	userRepo := userRepo.RepositoryFactory(dbCon)

	userService := userService.NewUserService(userRepo)
	authService := authService.NewAuthService(userRepo)
	jwtService := jwtService.NewJWTService()

	controller := api.Controller{
		Auth: auth.NewAuthController(authService, jwtService, userService),
		User: user.NewUserController(userService, jwtService),
	}

	return controller
}
