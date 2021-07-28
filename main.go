package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	dbuser   = "postgres"
	password = "postgres"
	dbname   = "test_db"
)

type user struct {
	UserID   string `json:"userID"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type registeredUserResponse struct {
	Status string `json:"status"`
}

type dbStore struct {
	db *sql.DB
}

type CookieStore struct {
}

var sessions = []CookieStore{}

func saveUserToDB(conn *dbStore, newUser user) error {
	var query = `INSERT INTO users(username, password) VALUES ($1, $2) RETURNING user_id;`
	_, err := conn.db.Exec(query, newUser.Username, newUser.Password)
	return err
}

func (conn *dbStore) registerUser(c *gin.Context) {
	var newUser user
	var response registeredUserResponse

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	var query = `INSERT INTO users(username, password) VALUES ($1, $2) RETURNING user_id;`
	_, err := conn.db.Exec(query, newUser.Username, newUser.Password)

	//err := saveUserToDB(conn, newUser)
	if err != nil {
		response.Status = "Server Error"
		c.IndentedJSON(http.StatusBadRequest, response)
		return
	}
	response.Status = "Account created!"

	var rows *sql.Rows
	rows, err = conn.db.Query("SELECT * FROM users WHERE username = $1 LIMIT 1", newUser.Username)
	checkError(err)

	var loggedUser = ""

	if rows.Next() {
		var userID string
		var username string
		var password string
		err = rows.Scan(&userID, &username, &password)
		checkError(err)

		loggedUser = userID
	}

	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: "sid", Value: loggedUser, Expires: expiration, HttpOnly: true}
	http.SetCookie(c.Writer, &cookie)
	c.IndentedJSON(http.StatusCreated, response)
}

func (conn *dbStore) connectToDB(connStr string) {
	var err error
	var db *sql.DB

	db, err = sql.Open("postgres", connStr)
	checkError(err)

	conn.db = db

	fmt.Printf("\nSuccessfully connected to database!\n")
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
	fmt.Println("Connection Tested, Success!")
}

func main() {
	router := gin.Default()
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable",
		dbuser,
		dbname,
		password,
		host)

	// *********** Store instance **************
	dbStoreI := dbStore{db: nil}

	// *********** Connecting to DB **************
	dbStoreI.connectToDB(connStr)
	defer dbStoreI.closeConnection()

	// *********** Testing connection **************
	dbStoreI.testConnection()

	//var rows *sql.Rows
	//rows, err := dbStoreI.db.Query("SELECT * FROM users")
	//checkError(err)
	//
	//for rows.Next() {
	//	var userID string
	//	var username string
	//	var password string
	//
	//	err = rows.Scan(&userID, &username, &password)
	//	checkError(err)
	//
	//	fmt.Println("userID | username | password")
	//	fmt.Printf("%s | %s | %s\n", userID, username, password)
	//}

	router.POST("/register", dbStoreI.registerUser)

	err := router.Run("localhost:8080")
	checkError(err)
}
