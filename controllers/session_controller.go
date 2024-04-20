package controllers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetSession(c *gin.Context) bool {
	// 获取当前用户的session，来检查是否登录
	session := sessions.Default(c)
	loginuser := session.Get("loginuser")
	fmt.Println("loginuser:", loginuser)
	if loginuser != nil {
		return true
	} else {
		return false
	}
}
