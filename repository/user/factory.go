package user

import (
	"rest-api/business/user"
	"rest-api/util"
)

func RepositoryFactory(dbCon *util.DatabaseConnection) user.UserRepository {
	var userRepository user.UserRepository

	if dbCon.Driver == util.MySQL {
		userRepository = NewMysqlRepository(dbCon.MySQL)
	} else if dbCon.Driver == util.PostgreSQL {
		userRepository = NewPostgresRepository(dbCon.PostgreSQL)
	} else {
		panic("Database driver not supported")
	}

	return userRepository
}
