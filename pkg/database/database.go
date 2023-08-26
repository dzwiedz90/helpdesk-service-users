package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/dzwiedz90/helpdesk-service-users/model"
	"github.com/dzwiedz90/helpdesk-service-users/service/serviceconfig"
)

func InsertUser(ctx context.Context, tx *sql.Tx, creq *model.CreateUser, cfg *serviceconfig.ServerConfig) (int64, error) {
	query := "INSERT INTO users (username, password, email, first_name, last_name, age, gender, street, city, postal_code, country, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)"
	_, err := tx.ExecContext(ctx, query, creq.Username, creq.Password, creq.Email, creq.FirstName, creq.LastName, creq.Age, creq.Gender, creq.Address.Street, creq.Address.City, creq.Address.PostalCode, creq.Address.Country, time.Now(), time.Now())
	if err != nil {
		cfg.Logger.ErrorLogger(fmt.Sprintf("Failed to insert user into DB: %v", err))
		return 0, err
	}

	query = "SELECT id FROM users ORDER BY created_at DESC LIMIT 1;"
	row := tx.QueryRow(query)

	var id int64

	err = row.Scan(&id)
	if err != nil {
		message := fmt.Sprintf("Failed getting user id from database: %v", err)
		cfg.Logger.ErrorLogger(message)
		return 0, err
	}

	return id, nil
}

func GetUser(ctx context.Context, tx *sql.Tx, id int64, cfg *serviceconfig.ServerConfig) (*model.User, error) {
	var user model.User
	var address model.Address

	query := "SELECT id, username, email, first_name, last_name, age, gender, street, city, postal_code, country FROM users WHERE id=$1"

	row := tx.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&user.UserId,
		&user.Username,
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
		cfg.Logger.ErrorLogger(fmt.Sprintf("Failed to get user from the database: %v", err))
		return nil, err
	}

	return &user, nil
}

func GetAllUsers(ctx context.Context, tx *sql.Tx, cfg *serviceconfig.ServerConfig) ([]*model.User, error) {
	var users []*model.User

	query := "SELECT id, username, email, first_name, last_name, age, gender, street, city, postal_code, country FROM users"

	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return users, nil
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		var address model.Address
		user.Address = address
		err := rows.Scan(
			&user.UserId,
			&user.Username,
			&user.Email,
			&user.FirstName,
			&user.LastName,
			&user.Age,
			&user.Gender,
			&user.Address.Street,
			&user.Address.City,
			&user.Address.PostalCode,
			&user.Address.Country,
		)

		if err != nil {
			return users, err
		}

		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		return users, err
	}

	return users, nil
}

func UpdateUser(ctx context.Context, tx *sql.Tx, creq *model.UpdateUser, cfg *serviceconfig.ServerConfig) error {
	query := `UPDATE users SET email=$1, first_name=$2, last_name=$3, age=$4, gender=$5, street=$6, city=$7, postal_code=$8, country=$9, updated_at=$10 WHERE id=$11`

	_, err := tx.ExecContext(
		ctx,
		query,
		creq.Email,
		creq.FirstName,
		creq.LastName,
		creq.Age,
		creq.Gender,
		creq.Address.Street,
		creq.Address.City,
		creq.Address.PostalCode,
		creq.Address.Country,
		time.Now(),
		creq.UserId,
	)

	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(ctx context.Context, tx *sql.Tx, id int64, cfg *serviceconfig.ServerConfig) error {
	query := `DELETE FROM users WHERE id=$1`

	_, err := tx.ExecContext(
		ctx,
		query,
		id,
	)

	if err != nil {
		return err
	}

	return nil
}
