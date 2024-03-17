package model

import "time"

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
}

type Movie struct {
	ID          string  `db:"id" json:"id"`
	Title       string  `db:"title" json:"title"`
	Rating      float32 `db:"rating" json:"rating"`
	ReleaseYear uint8   `db:"release_year" json:"release_year"`
}
