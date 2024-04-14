package models

import (
	"errors"
	"rest-api/database"
	"rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() (error) {
	query := `
	INSERT INTO users(email, password)
	VALUES(?, ?)`
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	u.ID = id

	return err
}

func (u *User) ValidateCredentials() (error) {
	query := `
	SELECT id, password
	FROM users
	WHERE email = ?`

	row := database.DB.QueryRow(query, u.Email)
	
	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if(err != nil) {
		return err
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if (!passwordIsValid) {
		return errors.New("invalid credentials")
	}
	return nil
}