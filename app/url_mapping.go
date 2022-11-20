package app

import (
	"microservices-bookstore_users/controllers/ping"
	"microservices-bookstore_users/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
	router.GET("/users/search", users.SearchUser)
}
