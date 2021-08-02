package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"restful.go/restapi/cookies"
	"restful.go/restapi/dbcalls"
	"restful.go/restapi/structs"
	"restful.go/restapi/utils"
)

const FIND_USER = "SELECT * FROM users WHERE username = $1 LIMIT 1"

type loginUser = structs.LoginUser
type loginUserResponse = structs.LoginUserResponse

func LoginHandler(c *gin.Context) {
	cookieStore := cookies.CookieStoreI
	var user loginUser
	var response loginUserResponse
	db, err := dbcalls.GetDB().DB, c.BindJSON(&user)
	utils.CheckError(err)

	row, err := db.Query(FIND_USER, user.Username)
	utils.CheckError(err)

	var loggedUser = ""

	if row.Next() {
		response.Status = "Logged In"
		response.Error = ""

		var userID string
		var username string
		var password string
		err = row.Scan(&userID, &username, &password)
		utils.CheckError(err)

		loggedUser = userID
		user.UserID = loggedUser

		cookie := utils.BuildCookie(loggedUser)
		utils.SetAndStoreCookieLogin(c, cookie, user, cookieStore)

		c.IndentedJSON(http.StatusCreated, response)
		return
	} else {
		response.Status = "Server Error"
		response.Error = "User doesn't exist"
		c.IndentedJSON(http.StatusCreated, response)
	}
}
