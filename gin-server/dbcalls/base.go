package dbcalls

import (
	"database/sql"
	"fmt"

	"restful.go/restapi/utils"
)

const (
	host     = "localhost"
	dbuser   = "postgres"
	password = "postgres"
	dbname   = "test_db"
)

type DBStore struct {
	DB *sql.DB
}

var DBStoreI = DBStore{DB: nil}

func GetDB() DBStore {
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable",
		dbuser,
		dbname,
		password,
		host)
	if DBStoreI.DB == nil {
		fmt.Println("Connecting to DB!")
		DBStoreI.ConnectToDB(connStr)
	}
	return DBStoreI
}

func (conn *DBStore) ConnectToDB(connStr string) {
	var err error
	var db *sql.DB

	db, err = sql.Open("postgres", connStr)
	utils.CheckError(err)

	conn.DB = db

	fmt.Printf("\nSuccessfully connected to database!\n")
}

func (conn DBStore) CloseConnection() {
	err := conn.DB.Close()
	utils.CheckError(err)
}

func (conn DBStore) TestConnection() {
	err := conn.DB.Ping()
	utils.CheckError(err)
	fmt.Println("Connection Tested, Success!")
}
