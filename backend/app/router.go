package app

import "github.com/lucasbarroso23/go-react-auth/backend/controller/users"

func mapUrls() {
	router.POST("/api/register", users.Register)
}
