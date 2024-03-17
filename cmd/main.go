package main

import (
	"FilmLibrary/configuration"
	"FilmLibrary/handler"
	"FilmLibrary/model"
	"FilmLibrary/storage"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"net/http"
	"strconv"
)

// Путь к файлу конфигурации.
const configPath = "configuration.json"

func main() {
	// Получаем конфиг.
	config := configuration.GetConfig(configPath)

	// Подключаемся к базе данных.
	dataBase, err := sqlx.Open("postgres", getDSN(config.DBConf))
	if err != nil {
		panic(err)
	}

	server := handler.Server{
		Storage: &storage.Storage{
			DB: dataBase,
		},
	}

	// Создаём мультиплексер или как его.
	router := http.NewServeMux()

	// HandleFunc обрабатывает URL-маршруты.
	router.HandleFunc("/actors", server.ActorsHandler)
	router.HandleFunc("/cast", server.CastHandler)

	port := ":" + strconv.Itoa(config.Port)

	// Запускаем сервер на прослушку порта 8080.
	err = http.ListenAndServe(port, router)
	if err != nil {
		panic(err)
	}
}

// getDSN функция создаёт строку Data Source Name.
func getDSN(cfg model.DBConf) string {
	return fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		cfg.User, cfg.Password, cfg.DBName, cfg.Sslmode)
}
