package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	dbuser   = "postgres"
	password = "postgres"
	dbname   = "test_db"
)

// type user struct {
// 	UserID   string `json:"userID"`
// 	Username string `json:"username"`
//}

// func getAlbums(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, albums)
// }

// func postAlbums(c *gin.Context) {
// 	var newAlbum album

// 	if err := c.BindJSON(&newAlbum); err != nil {
// 		return
// 	}

// 	albums = append(albums, newAlbum)
// 	c.IndentedJSON(http.StatusCreated, newAlbum)
// }

type dbStore struct {
	db *sql.DB
}

func (conn *dbStore) connectToDB(connStr string) *sql.DB {
	var err error
	var db = (*conn).db

	db, err = sql.Open("postgres", connStr)
	checkError(err)

	fmt.Printf("\nSuccessfully connected to database!\n")
	return db
}

func (conn dbStore) closeConnection() {
	err := conn.db.Close()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func (conn dbStore) testConnection() {
	err := conn.db.Ping()
	checkError(err)
}

func main() {
	router := gin.Default()
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable",
		dbuser,
		dbname,
		password,
		host)

	// ***********Store object**************
	dbStoreI := dbStore{ db: nil }

	dbStoreI.connectToDB(connStr)
	defer dbStoreI.closeConnection()

	dbStoreI.testConnection()

	//sqlStatement := `INSERT INTO students (name) VALUES ($1)`
	//_, err := db.Exec(sqlStatement, "Aditya")
	//if err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println("\nRow inserted successfully!")
	//}
	// router.GET("/albums", getAlbums)
	// router.POST("/albums", postAlbums)
	err := router.Run("localhost:8080")
	checkError(err)
}
