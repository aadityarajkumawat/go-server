package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"restful.go/restapi/dbcalls"
	"restful.go/restapi/handlers"
	"restful.go/restapi/utils"
)

func main() {
	router := gin.Default()

	// ************* Setting CORS Middleware ****************
	router.Use(utils.Cors())

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// *********** Connecting to DB **************
	db := dbcalls.GetDB()
	defer db.CloseConnection()

	// *********** Testing connection **************
	db.TestConnection()

	// ************ ROUTES **************
	router.POST("/register", handlers.RegisterHandler)
	router.POST("/login", handlers.LoginHandler)
	router.GET("/song", handlers.SongHandler)

	// *********** STARTING THE APP *************
	err := router.Run("localhost:8080")
	utils.CheckError(err)
}
