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
	"strconv"
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

	port := ":" + strconv.Itoa(config.Port)

	// Запускаем сервер на прослушку порта port (8080).
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
