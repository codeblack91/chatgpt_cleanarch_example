package repositories

import (
	"database/sql"

	"cleanarchtest/internal/models"
)

// UserRepository - реализация интерфейса UserRepository для работы с SQLite
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository - конструктор репозитория пользователей
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// GetAll - метод для получения всех пользователей из базы данных
func (r *UserRepository) GetAll() ([]models.User, error) {
	rows, err := r.db.Query("SELECT id, name, role FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Role); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// GetAll - метод для получения всех пользователей из базы данных
func (r *UserRepository) GetUserById(userID string) (models.User, error) {
	var user models.User
	rows, err := r.db.Query("SELECT id, name, role FROM users Where id = " + userID)
	if err != nil {
		return user, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Role); err != nil {
			return user, err
		}
		//users = append(users, user)
	}
	return user, nil
}
