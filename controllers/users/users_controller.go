package users

import (
	"microservices-bookstore_users/domain/users"
	"microservices-bookstore_users/services"
	"microservices-bookstore_users/utils/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	counter int
)

func getUserId(userIdParam string) (int64, *errors.RestError) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("user id should be a number")
	}
	return userId, nil
}

func Get(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		//TODO: Handle user creation error
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)

}

func Create(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		//TODO: Return bad request to the caller.
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//TODO: Handle user creation error
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func Update(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		//TODO: Return bad request to the caller.
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch

	result, updateErr := services.UpdateUser(isPartial, user)
	if updateErr != nil {
		//TODO: Handle user creation error
		c.JSON(updateErr.Status, updateErr)
		return
	}
	c.JSON(http.StatusOK, result)

}

func Delete(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	deleteErr := services.DeleteUser(userId)
	if deleteErr != nil {
		//TODO: Handle user creation error
		c.JSON(deleteErr.Status, deleteErr)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})

}

func Search(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")

}
