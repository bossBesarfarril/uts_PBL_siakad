package repository

import (
	"context"
	"uts-siakad/app/model"
	"uts-siakad/database"
)

func FindUserByEmail(email string) (model.User, error) {
	var user model.User
	query := "SELECT id, email, password, role FROM users WHERE email = $1"
	err := database.DB.QueryRow(context.Background(), query, email).Scan(&user.ID, &user.Email, &user.Password, &user.Role)
	return user, err
}

func FindUserByID(id int) (model.User, error) {
	var user model.User
	query := "SELECT id, email, password, role FROM users WHERE id = $1"
	err := database.DB.QueryRow(context.Background(), query, id).Scan(&user.ID, &user.Email, &user.Password, &user.Role)
	return user, err
}
