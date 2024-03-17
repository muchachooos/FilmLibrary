package main

import (
	"FilmLibrary/configuration"
	"FilmLibrary/handler"
	"FilmLibrary/model"
	"FilmLibrary/storage"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
)

// Путь к файлу конфигурации.
const configPath = "configuration.json"

func main() {
	// Получаем конфиг.
	config := configuration.GetConfig(configPath)

	// Подключаемся к базе данных.
	dataBase, err := sql.Open("postgres", getDSN(config.DBConf))
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

	// HandleFunc обрабатывает URL-маршрут (/actors), с помощью указанной функции (actorsHandler).
	router.HandleFunc("/actors", server.ActorsHandler)

	// Запускаем сервер на прослушку порта (8080).
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}

func getDSN(cfg model.DBConf) string {
	return fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		cfg.User, cfg.Password, cfg.DBName, cfg.Sslmode)
}
