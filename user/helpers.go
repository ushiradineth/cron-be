package user

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/ushiradineth/cron-be/auth"
)

func GetUser(idOrEmail string) (*User, error) {
	user := User{}
	var query string
	var args []interface{}

	id, err := uuid.Parse(idOrEmail)
	if err != nil {
		id = uuid.Nil
	}

	if id == uuid.Nil {
		query = "SELECT * FROM users WHERE email=$1"
		args = append(args, idOrEmail)
	} else {
		query = "SELECT * FROM users WHERE id=$1"
		args = append(args, id)
	}

	err = DB.Get(&user, query, args...)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func DoesUserExist(idStr string, email string) (bool, int, error) {
	var id uuid.UUID

	if idStr != "" {
		parsedID, err := uuid.Parse(idStr)

		if err != nil {
			return false, 0, err
		}

		id = parsedID
	}

	var userCount int
	var query string
	var args []interface{}

	if idStr != "" {
		query = "SELECT COUNT(*) FROM users WHERE id=$1 OR email=$2"
		args = append(args, id)
	} else {
		query = "SELECT COUNT(*) FROM users WHERE email=$1"
	}

	args = append(args, email)

	err := DB.Get(&userCount, query, args...)
	if err != nil {
		return false, 0, err
	}

	return userCount != 0, userCount, nil
}

func GetUserFromJWT(r *http.Request) (*User, error) {
	accessToken, err := auth.GetJWT(r)
	if err != nil {
		return nil, err
	}

	JWT := auth.ParseAccessToken(accessToken)

	if JWT.StandardClaims.Valid() != nil {
		return nil, errors.New(fmt.Sprint("Access Token is invalid or expired"))
	}

	user, err := GetUser(JWT.Id.String())
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%v", err))
	}

	return user, nil
}
