package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"restful.go/restapi/cookies"
	"restful.go/restapi/utils"
)

type song struct {
	Name   string `json:"name"`
	Length uint   `json:"length"`
	Artist string `json:"artist"`
}

func SongHandler(c *gin.Context) {
	cookieStore := cookies.CookieStoreI
	var oneSong = song{Name: "Golmaal", Artist: "Brijesh Shandilya", Length: 3}
	cookie, err := c.Request.Cookie("sid")
	utils.CheckError(err)
	if (*cookieStore).GetCookie(cookie.Value) != "" {
		c.IndentedJSON(http.StatusCreated, oneSong)
	} else {
		c.IndentedJSON(http.StatusCreated, song{})
	}
}
