package handler

import (
	"FilmLibrary/helpers/id"
	"FilmLibrary/helpers/parser"
	"FilmLibrary/model"
	"encoding/json"
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
			res.Write([]byte("No name"))
			return
		}

		actor := model.Actor{
			ID:     id.NewID(),
			Name:   actorReq.Name,
			Gender: actorReq.Gender,
			Birth:  actorReq.Birth,
		}

		err = s.Storage.CreateActor(actor)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			res.Write([]byte("Error: Create Actor error"))
			return
		}

		res.WriteHeader(http.StatusOK)
		return

	case http.MethodGet:
		actorId := req.FormValue("id")
		if actorId == "" {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte("No actor ID"))
			return
		}

		actor, err := s.Storage.GetActor(actorId)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			res.Write([]byte("Error: Get Actor error"))
			return
		}

		response, err := json.Marshal(actor)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			res.Write([]byte("Error: Marshal error"))
			return
		}

		res.WriteHeader(http.StatusOK)
		res.Write(response)

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
			res.Write([]byte("No name"))
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
			res.Write([]byte("Error: Marshal error"))
			return
		}

		res.WriteHeader(http.StatusOK)
		res.Write(response)

		return

	case http.MethodDelete:

	}
}
