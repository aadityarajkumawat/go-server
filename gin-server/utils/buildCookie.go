package utils

import (
	"net/http"
	"time"
)

func BuildCookie(loggedUser string) http.Cookie {
	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: "sid", Value: loggedUser, Expires: expiration, HttpOnly: true}
	return cookie
}
