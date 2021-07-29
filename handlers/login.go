package handlers

import (
	"github.com/gin-gonic/gin"
)

const FIND_USER = ""

type loginUser struct {
	UserID string `json:"userID"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginUserResponse struct {
	Status string `json:"status"`
	Error string `json:"error"`
}

func LoginHandler(c *gin.Context) {
	//cookieStore := cookies.CookieStoreI
	//var user loginUser
	//var response loginUserResponse
	//db, err := dbcalls.GetDB().DB, c.BindJSON(&user)
	//utils.CheckError(err)

	//c.Request.Body


	//buf := new(bytes.Buffer)
	//_, err := buf.ReadFrom(c.Request.Body)
	//if err != nil {
	//	return
	//}
	//newStr := buf.String()
	//fmt.Println(newStr)

	//_, err = db.Query()
}
