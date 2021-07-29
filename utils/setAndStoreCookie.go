package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type cookieI struct {
	UserSessionID string
}

var cookiesStore = make([]cookieI, 0, 10)

func SetAndStoreCookie(c *gin.Context, cookie http.Cookie, loggedUser string) {
	http.SetCookie(c.Writer, &cookie)
	storedCookie := cookieI{UserSessionID: loggedUser}
	cookiesStore = append(cookiesStore, storedCookie)
}
