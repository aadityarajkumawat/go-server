package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"restful.go/restapi/cookies"
	"restful.go/restapi/dbcalls"
	"restful.go/restapi/structs"
	"restful.go/restapi/utils"
	s "strings"
)

const INSERT_USER = "INSERT INTO users(username, password) VALUES ($1, $2) RETURNING user_id;"
const SELECT_LOGGED_USER = "SELECT * FROM users WHERE username = $1 LIMIT 1"

type registeredUserResponse = structs.RegisteredUserResponse
type registerUser = structs.RegisterUser

func RegisterHandler(c *gin.Context) {
	cookieStore := cookies.CookieStoreI
	var newUser registerUser
	var response registeredUserResponse
	db, err := dbcalls.GetDB().DB, c.BindJSON(&newUser)
	utils.CheckError(err)

	_, err = db.Exec(INSERT_USER, newUser.Username, newUser.Password)
	if err != nil {
		response.Status = "Server Error"
		errStr := err.Error()
		isDup := s.Contains(errStr, "username_key")
		if isDup {
			response.Error = "Username is not unique"
		}
		c.IndentedJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = "Account created!"
	response.Error = "--no-error--"

	rows, err := db.Query(SELECT_LOGGED_USER, newUser.Username)
	utils.CheckError(err)

	var loggedUser = ""

	if rows.Next() {
		var userID string
		var username string
		var password string
		err = rows.Scan(&userID, &username, &password)
		utils.CheckError(err)

		loggedUser = userID
		newUser.UserID = loggedUser
	}

	cookie := utils.BuildCookie(loggedUser)
	utils.SetAndStoreCookieRegister(c, cookie, newUser, cookieStore)

	c.IndentedJSON(http.StatusCreated, response)
}
