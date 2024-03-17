package handler

import (
	"FilmLibrary/helpers/idgen"
	"FilmLibrary/helpers/parser"
	"FilmLibrary/model"
	"encoding/json"
	"errors"
	"net/http"
)

func (s *Server) ActorsHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:

		var actorReq model.Actor
		err := parser.ParseBody(req, &actorReq)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		if actorReq.Name == "" {
			res.WriteHeader(http.StatusBadRequest)
			errors.New("no name")
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
			res.WriteHeader(http.StatusInternalServerError)
			errors.New("create Actor error")
			return
		}

		res.WriteHeader(http.StatusOK)
		return

	case http.MethodPut:
		var actorReq model.Actor
		err := parser.ParseBody(req, &actorReq)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		if actorReq.Name == "" {
			res.WriteHeader(http.StatusBadRequest)
			errors.New("no name")
			return
		}

		actor := model.Actor{
			ID:     actorReq.ID,
			Name:   actorReq.Name,
			Gender: actorReq.Gender,
			Birth:  actorReq.Birth,
		}

		actor, err = s.Storage.UpdateActor(actor)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			res.Write([]byte(err.Error()))
			return
		}

		response, err := json.Marshal(actor)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			errors.New("marshal error")
			return
		}

		res.WriteHeader(http.StatusOK)
		res.Write(response)
		return

	case http.MethodDelete:
		actorId := req.FormValue("idgen")
		if actorId == "" {
			res.WriteHeader(http.StatusBadRequest)
			errors.New("no actor ID")
			return
		}

		err := s.Storage.DeleteActor(actorId)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			errors.New("error: Get Actor error")
			return
		}

		res.WriteHeader(http.StatusOK)
		return
	}
}

func (s *Server) CastHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		actorId := req.FormValue("actor_id")
		if actorId == "" {
			res.WriteHeader(http.StatusBadRequest)
			errors.New("no actor ID")
			return
		}

		movieId := req.FormValue("movie_id")
		if movieId == "" {
			res.WriteHeader(http.StatusBadRequest)
			errors.New("no movie ID")
			return
		}

		err := s.Storage.AddInCast(actorId, movieId)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			errors.New("added to the cast error")
			return
		}

		res.WriteHeader(http.StatusOK)
		return

	case http.MethodDelete:
		actorId := req.FormValue("actor_id")
		if actorId == "" {
			res.WriteHeader(http.StatusBadRequest)
			errors.New("no actor ID")
			return
		}

		movieId := req.FormValue("movie_id")
		if movieId == "" {
			res.WriteHeader(http.StatusBadRequest)
			errors.New("no movie ID")
			return
		}

		err := s.Storage.DeleteFromCast(actorId, movieId)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			errors.New("delete cast error")
			return
		}

		res.WriteHeader(http.StatusOK)
		return
	}
}
