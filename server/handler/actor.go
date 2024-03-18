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

func (s *Server) ActorsHandler(res http.ResponseWriter, req *http.Request) {
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
	case http.MethodPost:
		s.createActor(res, req)
		return
	case http.MethodGet:
		s.getActorsWithMovies(res)
		return
	case http.MethodPut:
		s.deleteActors(res, req)
		return
	case http.MethodDelete:
		s.deleteActors(res, req)
		return
	default:
		render.NotFound(res)
	}
}

// createActor godoc
// @Tags Actors
// @Description Get movies
// @Param params body model.Actor true "movie"
// @Produce json
// @Success 200
// @Failure 500
// @Router /actors [post]
func (s *Server) createActor(res http.ResponseWriter, req *http.Request) {
	var actorReq model.Actor
	err := parser.ParseBody(req, &actorReq)
	if err != nil {
		render.BadRequest(res, err)
		return
	}

	if actorReq.Name == "" {
		render.BadRequest(res, errors.New("no name passed"))
		return
	}

	actor := model.Actor{
		ID:     idgen.NewID(),
		Name:   actorReq.Name,
		Gender: actorReq.Gender,
		Birth:  actorReq.Birth,
	}

	err = s.Storage.CreateActor(actor)
	if err != nil {
		render.Internal(res, err)
		return
	}

	render.OK(res)
	log.Print("create actor request completed successful")
}

// getActorsWithMovies godoc
// @Tags Actors
// @Description Get actors with movies
// @Produce json
// @Success 200 {array} model.Actor "Response body"
// @Failure 500
// @Router /actors [get]
func (s *Server) getActorsWithMovies(res http.ResponseWriter) {
	actors, err := s.Storage.GetActorsWithMovies()
	if err != nil {
		render.Internal(res, err)
		return
	}

	render.JSON(res, actors)
	log.Print("get actors request completed successful")
}

// deleteActors godoc
// @Tags Movies
// @Description delete actor
// @Produce json
// @Param id query string true "Actor ID"
// @Success 200
// @Failure 500
// @Router /movies [delete]
func (s *Server) deleteActors(res http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	if id == "" {
		render.BadRequest(res, errors.New("no id passed"))
	}

	err := s.Storage.DeleteActor(id)
	if err != nil {
		render.Internal(res, err)
		return
	}

	render.OK(res)
	log.Print("delete actor request completed successful")
}

// updateActors godoc
// @Tags Actors
// @Description Update actors
// @Param params body model.Actor true "movie"
// @Produce json
// @Success 200
// @Failure 500
// @Router /actors [put]
func (s *Server) updateActors(res http.ResponseWriter, req *http.Request) {
	var actor model.Actor
	err := parser.ParseBody(req, &actor)
	if err != nil {
		render.BadRequest(res, err)
		return
	}

	err = s.Storage.UpdateActor(actor)
	if err != nil {
		render.Internal(res, err)
		return
	}

	render.OK(res)
	log.Print("update actor request completed successful")
}
