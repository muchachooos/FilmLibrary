package storage

import (
	"database/sql"
	"errors"
	"filmlib/helpers/auth"
	"filmlib/model"

	"github.com/google/uuid"
)

func (s *Storage) Authorize(log, pass string) (string, error) {
	var hashedPass string
	err := s.DB.Get(&hashedPass, "SELECT hashed_pass FROM user_auth WHERE login = $1", log)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", model.ErrorAuth
		}
		return "", err
	}

	err = auth.CompareHashPassword(hashedPass, pass)
	if err != nil {
		return "", model.ErrorAuth
	}

	token := uuid.NewString()

	_, err = s.DB.Exec("UPDATE user_auth SET token = $1 WHERE login = $2 AND hashed_pass = $3", token, log, hashedPass)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Storage) CheckToken(token string) (model.User, error) {
	var user model.User
	err := s.DB.Get(&user, "SELECT * FROM user_auth WHERE token = $1", token)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, model.ErrorAuth
		}
		return model.User{}, err
	}

	return user, nil
}
