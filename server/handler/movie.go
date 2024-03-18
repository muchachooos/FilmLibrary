package handler

import (
	"errors"
	"filmlib/helpers/idgen"
	"filmlib/helpers/parser"
	"filmlib/helpers/render"
	"filmlib/model"
	"log"
	"net/http"
)

func (s *Server) Movies(res http.ResponseWriter, req *http.Request) {
	user, err := s.checkAuth(req)
	if err != nil {
		if errors.Is(err, model.ErrorAuth) {
			render.Unauthorized(res)
			return
		}
		render.Internal(res, err)
		return
	}

	// только админ имет право на редактирование данных
	if req.Method != http.MethodGet && !user.IsAdmin {
		render.NotAllowed(res)
		return
	}

	switch req.Method {
	case http.MethodGet:
		s.getMovies(res, req)
		return
	case http.MethodPost:
		s.createMovie(res, req)
		return
	case http.MethodPut:
		s.updateMovie(res, req)
		return
	case http.MethodDelete:
		s.deleteMovie(res, req)
		return
	default:
		render.NotFound(res)
	}
}

// SearchMoviesByTitle godoc
// @Tags Movies
// @Summary Get movies
// @Description Get movies
// @Param partOfTitle query string true "part of movie title"
// @Produce json
// @Success 200 {array} model.Movie "Response body"
// @Failure 500
// @Router /movies/search-by-title [get]
func (s *Server) SearchMoviesByTitle(w http.ResponseWriter, req *http.Request) {
	_, err := s.checkAuth(req)
	if err != nil {
		if errors.Is(err, model.ErrorAuth) {
			render.Unauthorized(w)
			return
		}
		render.Internal(w, err)
		return
	}

	partOfTitle := req.FormValue("partOfTitle")
	if partOfTitle == "" {
		render.BadRequest(w, errors.New("no partOfTitle passed"))
		return
	}

	movies, err := s.Storage.SearchMoviesByTitle(partOfTitle)
	if err != nil {
		render.Internal(w, err)
		return
	}

	render.JSON(w, movies)
}

// SearchMoviesByActor godoc
// @Tags Movies
// @Summary Get movies
// @Description Get movies
// @Param partOfActor query string true "part of actor's name"
// @Produce json
// @Success 200 {array} model.Movie "Response body"
// @Failure 500
// @Router /movies/search-by-actor [get]
func (s *Server) SearchMoviesByActor(w http.ResponseWriter, req *http.Request) {
	_, err := s.checkAuth(req)
	if err != nil {
		if errors.Is(err, model.ErrorAuth) {
			render.Unauthorized(w)
			return
		}
		render.Internal(w, err)
		return
	}

	partOfActor := req.FormValue("partOfActor")
	if partOfActor == "" {
		render.BadRequest(w, errors.New("no partOfActor query param"))
		return
	}

	movies, err := s.Storage.SearchMoviesByActor(partOfActor)
	if err != nil {
		render.Internal(w, err)
		return
	}

	render.JSON(w, movies)
}

// getMovies godoc
// @Tags Movies
// @Summary Get movies
// @Description Get movies
// @Param sortBy query string true "field to sort by"
// @Param sortType query string true "ASC DESC"
// @Produce json
// @Success 200 {array} model.Movie "Response body"
// @Failure 500
// @Router /movies [get]
func (s *Server) getMovies(res http.ResponseWriter, req *http.Request) {
	sortBy := req.FormValue("sortBy")
	sortType := req.FormValue("sortType")

	movies, err := s.Storage.GetMovies(sortBy, sortType)
	if err != nil {
		render.Internal(res, err)
		return
	}

	render.JSON(res, movies)
	log.Print("get movies request completed successful")
}

// deleteMovie godoc
// @Tags Movies
// @Description delete movie
// @Produce json
// @Param id query string true "Movie ID"
// @Success 200
// @Failure 500
// @Router /movies [delete]
func (s *Server) deleteMovie(res http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	if id == "" {
		render.BadRequest(res, errors.New("no id passed"))
	}

	err := s.Storage.DeleteMovie(id)
	if err != nil {
		render.Internal(res, err)
		return
	}

	render.OK(res)
	log.Print("delete movie request completed successful")
}

// createMovie godoc
// @Tags Movies
// @Description Create movie
// @Param params body model.Movie true "movie"
// @Produce json
// @Accept json
// @Success 200
// @Failure 500
// @Router /movies [post]
func (s *Server) createMovie(res http.ResponseWriter, req *http.Request) {
	var movie model.Movie
	err := parser.ParseBody(req, movie)
	if err != nil {
		render.BadRequest(res, err)
		return
	}

	movie.ID = idgen.NewID()

	err = s.Storage.CreateMovie(movie)
	if err != nil {
		render.Internal(res, err)
		return
	}

	render.OK(res)
	log.Print("create movie request completed successful")
}

// updateMovie godoc
// @Tags Movies
// @Description Update movie
// @Param params body model.Movie true "movie"
// @Produce json
// @Accept json
// @Success 200
// @Failure 500
// @Router /movies [put]
func (s *Server) updateMovie(res http.ResponseWriter, req *http.Request) {
	var movie model.Movie
	err := parser.ParseBody(req, movie)
	if err != nil {
		render.BadRequest(res, err)
		return
	}

	err = s.Storage.UpdateMovie(movie)
	if err != nil {
		render.Internal(res, err)
		return
	}

	render.OK(res)
	log.Print("update movie request completed successful")
}
