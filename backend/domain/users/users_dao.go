package users

import (
	"github.com/lucasbarroso23/go-react-auth/backend/datasource/mysql/users_db"
	"github.com/lucasbarroso23/go-react-auth/backend/utils/errors"
)

var (
	queryInsertUser = "INSERT INTO users (first_name, last_name, email, password) VALUES (?, ?, ?, ?);"
)

func (user *User) Save() *errors.RestErr {
	stat, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError("database error")
	}
	defer stat.Close()

	insertResult, err := stat.Exec(user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return errors.NewInternalServerError("database error")
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError("database error")
	}

	user.Id = userID
	return nil
}
