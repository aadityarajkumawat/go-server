package utils

import (
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func Cors() gin.HandlerFunc {
	var handleCorsFunc = cors.Middleware(cors.Config{
		Origins:         "http://localhost:3000",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		Credentials:     true,
		ValidateHeaders: false,
	})

	return handleCorsFunc
}
