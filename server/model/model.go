package model

import (
	"errors"
	"time"
)

type Err struct {
	Error string `json:"error"`
}

type Config struct {
	Port   int    `json:"port"`
	DBConf DBConf `json:"DataBase"`
}

type DBConf struct {
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dataBaseName"`
	Sslmode  string `json:"sslmode"`
}

type Actor struct {
	ID     string    `db:"id" json:"id"`
	Name   string    `db:"name" json:"name"`
	Gender string    `db:"gender" json:"gender"`
	Birth  time.Time `db:"date_of_birth" json:"date_of_birth"`

	Movies []Movie `db:"-" json:"movies"`
}

type Movie struct {
	ID          string  `db:"id" json:"id"`
	Title       string  `db:"title" json:"title"`
	Rating      float32 `db:"rating" json:"rating"`
	ReleaseYear uint8   `db:"release_year" json:"release_year"`
}

type CastRecord struct {
	ActorID string `db:"actor_id"`
	MovieID string `db:"movie_id"`
}

type User struct {
	ID         string `db:"id"`
	Login      string `db:"login"`
	HashedPass string `db:"hashed_pass"`
	Token      string `db:"token"`
	IsAdmin    bool   `db:"is_admin"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

type Request struct {
	Login string `json:"login"`
	Pass  string `json:"password"`
}

var ErrorAuth = errors.New("authorization unsuccessful")
