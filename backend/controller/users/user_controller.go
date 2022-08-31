package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasbarroso23/go-react-auth/backend/domain/users"
	"github.com/lucasbarroso23/go-react-auth/backend/services"
	"github.com/lucasbarroso23/go-react-auth/backend/utils/errors"
)

func Register(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		err := errors.NewBadRequestError("invalid json body")
		c.JSON(err.Status, err)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusOK, result)
}
