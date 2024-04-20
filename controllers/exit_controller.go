package controllers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ExitGet(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("loginuser")
	session.Save()
	fmt.Println("delete session...", session.Get("loginuser"))
	// 清除session后，重定向到/上
	c.Redirect(http.StatusMovedPermanently, "/")
}
