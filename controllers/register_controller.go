package controllers

import (
	"fmt"
	"gindemo/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterGet(c *gin.Context) {
	//返回html
	c.HTML(http.StatusOK, "register.html", gin.H{"title": "注册页"})
}

//处理注册

func RegisterPost(c *gin.Context) {
	//获取表单信息
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")
	repassword := c.PostForm("repassword")
	fmt.Println(username, email, password, repassword)

	//注册之前先判断该用户名是否已经被注册，如果已经注册，返回错误
	id := models.QueryUserWithEmail(email)
	fmt.Println("id:", id)
	if id > 0 {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "该邮箱已经被注册"})
		return
	}

	user := models.User{username, email, password}
	_, err := models.InsertUser(user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "注册失败"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "注册成功"})
	}
}
