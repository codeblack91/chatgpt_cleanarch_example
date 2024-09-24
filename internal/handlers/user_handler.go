package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "cleanarchtest/cmd/web/docs" // Path to your docs folder
	"cleanarchtest/internal/services"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func GetUsersHandler(userService *services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := userService.GetAllUsers()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Info("Получение всех пользователей")
		url := r.URL.String()
		fmt.Println("URL запроса:", url)

		// Возвращаем JSON-ответ
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}

// GetUserById - обработчик для получения списка пользователей
func GetUserByIdHandler(userService *services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		userID := vars["userID"]

		users, err := userService.GetUserById(userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Info("Получение конкретного пользователя")

		// Возвращаем JSON-ответ
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}
