package movie

import (
	"rest-api/business/movie"
	"rest-api/util"
)

func MovieRepositoryFactory(dbCon *util.DatabaseConnection) movie.MovieRepository {
	var movieRepository movie.MovieRepository

	if dbCon.Driver == util.MySQL {
		movieRepository = NewMysqlMovieRepository(dbCon.MySQL)
	} else if dbCon.Driver == util.PostgreSQL {
		movieRepository = NewPosgresMovieRepository(dbCon.PostgreSQL)
	} else {
		panic("Database driver not supported")
	}

	return movieRepository
}
