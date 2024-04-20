package controllers

import (
	"fmt"
	"gindemo/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginGet(c *gin.Context) {
	//返回html
	c.HTML(http.StatusOK, "login.html", gin.H{"title": "登录页"})
}

func LoginPost(c *gin.Context) {
	//获取表单信息
	email := c.PostForm("email")
	password := c.PostForm("password")
	fmt.Println("email:", email, ",password:", password)

	id := models.QueryUserWithParam(email, password)
	fmt.Println("id:", id)
	if id > 0 {
		// 将信息储存在cookie，浏览器发送请求时会自动带上cookie
		// 服务器端存储的是session，客户端存储的才是cookie
		// 通过cookie来判断这个用户是谁
		session := sessions.Default(c)
		// 将用户的登录邮箱存储到会话中，用于标识用户身份或者其他相关信息
		// loginuser 是存储在会话中的键，email 是要存储的值
		session.Set("loginuser", email)
		session.Save()

		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "登录成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "登录失败"})
	}
}
