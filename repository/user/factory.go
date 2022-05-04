package user

import (
	"rest-api/util"
)

func RepositoryFactory(dbCon *util.DatabaseConnection) UserRepository {
	var userRepository UserRepository

	if dbCon.Driver == util.MySQL {
		userRepository = NewUserRepo(dbCon.MySQL)
	}
	return userRepository
}
