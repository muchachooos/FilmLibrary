package model

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

type CreateActorReq struct {
	Name string `json:"name"`
}

type Actor struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}
