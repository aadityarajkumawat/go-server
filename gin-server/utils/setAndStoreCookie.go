package utils

import (
	"net/http"

	"restful.go/restapi/cookies"
	"restful.go/restapi/structs"

	"github.com/gin-gonic/gin"
)

type registerUser = structs.RegisterUser
type loginUser = structs.LoginUser

func SetAndStoreCookieRegister(c *gin.Context, cookie http.Cookie, loggedUser registerUser, cookieStore *cookies.CookieStore) {
	http.SetCookie(c.Writer, &cookie)
	storedCookie := cookies.CookieI{UserSessionID: loggedUser.UserID, Username: loggedUser.Username, Password: loggedUser.Password}
	cookieStore.SetCookie(storedCookie)
}

func SetAndStoreCookieLogin(c *gin.Context, cookie http.Cookie, loggedUser loginUser, cookieStore *cookies.CookieStore) {
	http.SetCookie(c.Writer, &cookie)
	storedCookie := cookies.CookieI{UserSessionID: loggedUser.UserID, Username: loggedUser.Username, Password: loggedUser.Password}
	cookieStore.SetCookie(storedCookie)
}
