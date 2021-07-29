package utils

import (
	"net/http"
	"restful.go/restapi/cookies"
	"restful.go/restapi/structs"

	"github.com/gin-gonic/gin"
)

type registerUser = structs.RegisterUser

func SetAndStoreCookie(c *gin.Context, cookie http.Cookie, loggedUser registerUser, cookieStore *cookies.CookieStore) {
	http.SetCookie(c.Writer, &cookie)
	storedCookie := cookies.CookieI{UserSessionID: loggedUser.UserID, Username: loggedUser.Username, Password: loggedUser.Password}
	cookieStore.SetCookie(storedCookie)
}
