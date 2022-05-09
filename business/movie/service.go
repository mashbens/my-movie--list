package movie

import (
	// "rest-api/business/movie/dto"

	"rest-api/business/movie/dto"
	"rest-api/business/movie/entity"
	_movie "rest-api/business/movie/response"

	"github.com/mashingan/smapping"
)

type MovieRepository interface {
	All(userID string) ([]entity.Movie, error)
	InsertMovie(movie entity.Movie) (entity.Movie, error)
	// DeleteMovie(movieID string, userID string) error
}

type MovieService interface {
	All(userID string) (*[]_movie.MovieResponse, error)
	AddMovie(addMovie dto.CreateMovieRequest, userID int64) (_movie.MovieResponse, error)
	// DeleteMovie(movieID string, userID string) error
}

type movieService struct {
	movieRepo MovieRepository
}

func NewMovieService(movieRepo MovieRepository) MovieService {
	return &movieService{
		movieRepo: movieRepo,
	}
}

func (c *movieService) All(userID string) (*[]_movie.MovieResponse, error) {
	movies, err := c.movieRepo.All(userID)
	if err != nil {
		return nil, err
	}
	movs := _movie.NewMovieArrayResponse(movies)
	return &movs, nil
}

func (c *movieService) AddMovie(addMovie dto.CreateMovieRequest, userID int64) (_movie.MovieResponse, error) {

	movie := entity.Movie{}
	err := smapping.FillStruct(&movie, smapping.MapFields(&addMovie))

	if err != nil {
		return _movie.MovieResponse{}, err
	}
	id := userID
	movie.UserID = id
	m, err := c.movieRepo.InsertMovie(movie)
	if err != nil {
		return _movie.MovieResponse{}, err
	}
	res := _movie.NewMovieResponse(m)
	return res, nil

}

// func (c *movieService) DeleteMovie(movieID string, userID string) error {
// 	return nil
// }
