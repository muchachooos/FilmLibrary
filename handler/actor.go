package handler

import (
	"FilmLibrary/helpers/id"
	"FilmLibrary/helpers/parser"
	"FilmLibrary/model"
	"fmt"
	"net/http"
)

func (s *Server) ActorsHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:

		var actorReq model.ActorJSON
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

		err = s.Storage.ActorsHandlerInDB(actor)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			res.Write([]byte("Error: Database error"))
			return
		}

		res.WriteHeader(http.StatusOK)
		return

	case http.MethodGet:
		id := req.FormValue("id")
		if id == "" {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte("No actor ID"))
			return
		}

		fmt.Println(id)

		res.WriteHeader(http.StatusOK)
		return
	}
}
