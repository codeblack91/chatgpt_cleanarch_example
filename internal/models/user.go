package models

type RoleEnum int

const (
	Admin RoleEnum = 1
	Users RoleEnum = 2
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role int    `json:"role"`
}
