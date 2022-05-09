package movie

import (
	"rest-api/business/movie"
	"rest-api/business/movie/entity"

	"gorm.io/gorm"
)

type MysqlmovieRepository struct {
	db *gorm.DB
}

func NewMysqlMovieRepository(db *gorm.DB) movie.MovieRepository {
	return &MysqlmovieRepository{
		db: db,
	}
}

func (c *MysqlmovieRepository) All(userID string) ([]entity.Movie, error) {
	movies := []entity.Movie{}

	c.db.Preload("User").Where("user_id = ?", userID).Find(&movies)
	return movies, nil
}

func (c *MysqlmovieRepository) InsertMovie(movie entity.Movie) (entity.Movie, error) {
	c.db.Create(&movie)
	c.db.Preload("User").Find(&movie)
	return movie, nil
}

// func (c *movieRepository) DeleteMovie(movieID string, userID string) error {
// 	var movie entity.Movie
// 	c.db.Where("id = ? AND user_id = ?", movieID, userID).First(&movie)
// 	if movie.ID == "" {
// 		return nil
// 	}
// 	c.db.Delete(&movie)
// 	return nil
// }
