package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"

	"cleanarchtest/internal/handlers"
	"cleanarchtest/internal/repositories"
	"cleanarchtest/internal/services"

	_ "cleanarchtest/cmd/web/docs" // Path to your docs folder

	"cleanarchtest/config"
)

var log = logrus.New()

// @title Example API
// @version 1.0
// @description This is a sample server.
// @host localhost:8080
// @BasePath /
func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	// Подключение к базе данных SQLite
	db, err := sql.Open(cfg.Database.Driver, cfg.Database.Connection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Настройка репозиториев и сервисов
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)

	// Настройка маршрутов
	router := mux.NewRouter()
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	router.HandleFunc("/users", handlers.GetUsersHandler(userService)).Methods("GET")
	router.HandleFunc("/users/{userID:[0-9]+}", handlers.GetUserByIdHandler(userService)).Methods("GET")

	// Запуск сервера
	//log.Println("Server started at :8080")
	//http.ListenAndServe(":8080", router)

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port),
		Handler: router,
	}

	log.Printf("Запуск сервера на %s:%d", cfg.Server.Host, cfg.Server.Port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
