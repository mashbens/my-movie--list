package movie

import (
	"fmt"
	"net/http"
	"rest-api/api/common/obj"
	"rest-api/api/common/response"
	movieService "rest-api/business/movie"
	"rest-api/business/movie/dto"
	jwtService "rest-api/business/user"
	"strconv"

	"github.com/eefret/gomdb"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type MovieController struct {
	movieService movieService.MovieService
	jwtService   jwtService.JWTService
}

func NewMovieController(
	movieService movieService.MovieService,
	jwtService jwtService.JWTService,
) *MovieController {
	return &MovieController{
		movieService: movieService,
		jwtService:   jwtService,
	}
}

func (controller *MovieController) All(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	token := controller.jwtService.ValidateToken(authHeader, c)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	movies, err := controller.movieService.All(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BuildErrorResponse("Failed to process request", "Failed to get movies", nil))
	}
	response := movies
	return c.JSON(http.StatusOK, response)
}
func (controller *MovieController) AddWishList(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	token := controller.jwtService.ValidateToken(authHeader, c)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	req := new(dto.MovieRequest)
	c.Bind(req)
	id := req.ID

	api := gomdb.Init("11ab3263")
	res, err := api.MovieByImdbID(id)
	if err != nil {
		return err
	}
	intuserID, _ := strconv.Atoi(userID)
	movie := dto.CreateMovieRequest{
		MovieID:  req.ID,
		Title:    res.Title,
		Year:     res.Year,
		Runtime:  res.Runtime,
		Released: res.Released,
		Genre:    res.Genre,
		Director: res.Director,
		Writer:   res.Writer,
		Actors:   res.Actors,
		Plot:     res.Plot,
		Language: res.Language,
		Country:  res.Country,
		Awards:   res.Awards,
		Poster:   res.Poster,
		UserID:   int64(intuserID),
	}

	movies, err := controller.movieService.AddMovie(movie, int64(intuserID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BuildErrorResponse("Failed to process request", "Failed to get movies", nil))
	}
	response := movies
	return c.JSON(http.StatusOK, response)
}

func (controller *MovieController) SearchMovie(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	token := controller.jwtService.ValidateToken(authHeader, c)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	_ = userID
	req := new(dto.MovieRequest)
	c.Bind(req)
	api := gomdb.Init("11ab3263")
	query := &gomdb.QueryData{Title: req.Search, SearchType: gomdb.MovieSearch}
	res, err := api.Search(query)
	if err != nil {
		fmt.Println(err)
		return err
	}
	response := res
	return c.JSON(http.StatusOK, response)
}

func (controller *MovieController) FindOneMovieByID(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	token := controller.jwtService.ValidateToken(authHeader, c)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	_ = userID

	req := new(dto.MovieRequest)
	c.Bind(req)

	reqId := req.ID
	movie, err := controller.movieService.FindMovieByID(reqId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BuildErrorResponse("Failed to process request", "Failed to get movies", nil))
	}
	response := movie
	return c.JSON(http.StatusOK, response)
}

func (controller *MovieController) DeleteMovie(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	token := controller.jwtService.ValidateToken(authHeader, c)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	req := new(dto.MovieRequest)
	c.Bind(req)

	reqId := req.ID

	err := controller.movieService.DeleteMovie(reqId, userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BuildErrorResponse("Failed to process request", "Failed to get movies", nil))
	}

	response := response.BuildResponse(true, "Movie deleted", obj.EmptyObj{})
	return c.JSON(http.StatusOK, response)
}
