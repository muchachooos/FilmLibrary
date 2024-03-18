package handler

import (
	"errors"
	"filmlib/helpers/render"
	"filmlib/model"
	"net/http"
)

func (s *Server) Auth(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		s.auth(res, req)
		return
	default:
		render.NotFound(res)
	}
}

// auth godoc
// @Tags Auth
// @Description Get auth token
// @Param login query string true "Login"
// @Param password query string true "password"
// @Produce json
// @Success 200 {object} model.AuthResponse
// @Failure 500
// @Router /movies [get]
func (s *Server) auth(res http.ResponseWriter, req *http.Request) {
	login := req.FormValue("login")
	if login == "" {
		render.BadRequest(res, errors.New("no login passed"))
	}
	password := req.FormValue("password")
	if password == "" {
		render.BadRequest(res, errors.New("no password passed"))
	}

	token, err := s.Storage.Authorize(login, password)
	if err != nil {
		render.Internal(res, err)
		return
	}

	resp := model.AuthResponse{
		Token: token,
	}

	render.JSON(res, resp)
}

func (s *Server) checkAuth(req *http.Request) (model.User, error) {
	authToken := req.Header.Get("Authorization")

	if authToken == "" {
		return model.User{}, errors.New("no Authorization Header")
	}

	return s.Storage.CheckToken(authToken)
}
