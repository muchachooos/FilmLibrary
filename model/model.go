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

type ActorJSON struct {
	Name   string    `json:"actor_name"`
	Gender string    `json:"gender"`
	Birth  time.Time `json:"date_of_birth"`
}

type Actor struct {
	ID     string    `db:"id"`
	Name   string    `db:"actor_name"`
	Gender string    `db:"gender"`
	Birth  time.Time `db:"date_of_birth"`
}
