package app

import (
	"microservices-bookstore_users/controllers/ping"
	"microservices-bookstore_users/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.Get)
	router.GET("/users/search", users.Search)
	router.POST("/users", users.Create)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
}
