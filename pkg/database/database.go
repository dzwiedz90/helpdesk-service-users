package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/dzwiedz90/helpdesk-service-users/logs"
	"github.com/dzwiedz90/helpdesk-service-users/model"
)

func InsertUser(ctx context.Context, tx *sql.Tx, creq *model.CreateUser) (int64, error) {
	query := "INSERT INTO users (username, password, email, first_name, last_name, age, gender, street, city, postal_code, country, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)"
	_, err := tx.ExecContext(ctx, query, creq.Username, creq.Password, creq.Email, creq.FirstName, creq.LastName, creq.Age, creq.Gender, creq.Address.Street, creq.Address.City, creq.Address.PostalCode, creq.Address.Country, time.Now(), time.Now())
	if err != nil {
		logs.ErrorLogger(fmt.Sprintf("Failed to insert user into DB: %v", err))
		return 0, err
	}

	query = "SELECT id FROM users ORDER BY created_at DESC LIMIT 1;"
	row := tx.QueryRow(query)

	var id int64

	err = row.Scan(&id)
	if err != nil {
		message := fmt.Sprintf("Failed getting user id from database: %v", err)
		logs.ErrorLogger(message)
		return 0, err
	}

	return id, nil
}

func GetUser(ctx context.Context, tx *sql.Tx, id int64) (*model.User, error) {
	var user model.User
	var address model.Address

	query := "SELECT id, username, password, email, first_name, last_name, age, gender, street, city, postal_code, country FROM users WHERE id=$1"

	row := tx.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&user.UserId,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Age,
		&user.Gender,
		&address.Street,
		&address.City,
		&address.PostalCode,
		&address.Country,
	)
	user.Address = address

	if err != nil {
		logs.ErrorLogger(fmt.Sprintf("Failed to get user from the database: %v", err))
		return nil, err
	}

	return &user, nil
}
