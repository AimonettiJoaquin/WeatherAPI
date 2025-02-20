package model

import "database/sql"

type User struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	OptOut           bool   `json:"optout"`
	NotificationTime string `json:"notification_time"`
}

func GetUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query(`SELECT id, name, email, optout, notification_time FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.OptOut, &user.NotificationTime); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func CreateUser(db *sql.DB, user *User) error {
	query := "INSERT INTO users (name, email, password, notification_time) VALUES (?, ?, ?, ?)"
	result, err := db.Exec(query, user.Name, user.Email, user.Password, user.NotificationTime)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = int(id)

	return nil
}

func GetUserByID(db *sql.DB, id int) (*User, error) {
	var user User
	query := "SELECT id, name, email, optout, notification_time FROM users WHERE id = ?"
	err := db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.OptOut, &user.NotificationTime)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateUser(db *sql.DB, user *User) error {
	query := "UPDATE users SET name = ?, email = ?, optout = ?, notification_time = ? WHERE id = ?"
	_, err := db.Exec(query, user.Name, user.Email, user.OptOut, user.NotificationTime, user.ID)
	if err != nil {
		return err
	}
	return nil
}
