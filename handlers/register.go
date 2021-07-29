package handlers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"restful.go/restapi/dbcalls"
	"restful.go/restapi/utils"
	s "strings"
)

type registerUser struct {
	UserID   string `json:"userID"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type registeredUserResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

func RegisterHandler(c *gin.Context) {
	var newUser registerUser
	var response registeredUserResponse
	var db = dbcalls.GetDB().DB

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	var query = "INSERT INTO users(username, password) VALUES ($1, $2) RETURNING user_id;"
	_, err := db.Exec(query, newUser.Username, newUser.Password)

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

	var rows *sql.Rows
	rows, err = db.Query("SELECT * FROM users WHERE username = $1 LIMIT 1", newUser.Username)
	utils.CheckError(err)

	var loggedUser = ""

	if rows.Next() {
		var userID string
		var username string
		var password string
		err = rows.Scan(&userID, &username, &password)
		utils.CheckError(err)

		loggedUser = userID
	}

	cookie := utils.BuildCookie(loggedUser)
	utils.SetAndStoreCookie(c, cookie, loggedUser)

	c.IndentedJSON(http.StatusCreated, response)
}
