package main

import (
	"filmlib/configuration"
	"filmlib/handler"
	"filmlib/model"
	"filmlib/storage"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Путь к файлу конфигурации.
const configPath = "configuration.json"

func main() {
	// Получаем конфиг.
	config := configuration.GetConfig(configPath)

	log.Printf("Config: %v", config)

	// Подключаемся к базе данных.
	dataBase, err := sqlx.Open("postgres", getDSN(config.DBConf))
	if err != nil {
		panic(err)
	}

	log.Print("Database connected")

	server := handler.Server{
		Config: config,
		Storage: &storage.Storage{
			DB: dataBase,
		},
	}

	runServer(server)
}

// @title FilmLibrary API
// @version 1.0

// @BasePath /
func runServer(server handler.Server) {
	// Создаём мультиплексер или как его.
	router := http.NewServeMux()

	// HandleFunc обрабатывает URL-маршруты
	router.HandleFunc("/actors", server.ActorsHandler)
	router.HandleFunc("/movies", server.Movies)
	router.HandleFunc("/movies/search-by-title", server.SearchMoviesByTitle)
	router.HandleFunc("/movies/search-by-actor", server.SearchMoviesByActor)

	// аутентификация
	router.HandleFunc("/auth", server.Auth)

	port := ":" + strconv.Itoa(server.Config.Port)

	log.Print("Server running")

	// Запускаем сервер на прослушку порта 8080
	err := http.ListenAndServe(port, router)
	if err != nil {
		panic(err)
	}
}

// getDSN функция создаёт строку Data Source Name.
func getDSN(cfg model.DBConf) string {
	return fmt.Sprintf("host=postgres port=5432 user=%s password=%s dbname=%s sslmode=%s",
		cfg.User, cfg.Password, cfg.DBName, cfg.Sslmode)
}
