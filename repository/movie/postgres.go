package movie

import (
	"rest-api/business/movie"
	"rest-api/business/movie/entity"

	"gorm.io/gorm"
)

type PostgreMovieRepository struct {
	db *gorm.DB
}

func NewPosgresMovieRepository(db *gorm.DB) movie.MovieRepository {
	return &PostgreMovieRepository{
		db: db,
	}
}

func (c *PostgreMovieRepository) All(userID string) ([]entity.Movie, error) {
	movies := []entity.Movie{}

	c.db.Preload("User").Where("user_id = ?", userID).Take(&movies)
	return movies, nil
}

func (c *PostgreMovieRepository) InsertMovie(movie entity.Movie) (entity.Movie, error) {
	c.db.Save(&movie)
	c.db.Preload("User").Find(&movie)
	return movie, nil
}
