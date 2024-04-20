package models

import (
	"fmt"
	"gindemo/database"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string
	Email    string
	Password string
}

// 插入用户
func InsertUser(user User) (int64, error) {
	return database.ModifyDB("insert into users(username,email,password) values (?,?,?)",
		user.Username, user.Email, user.Password)
}

// 按条件查询用户id
func QueryUserWightCon(con string) int {
	sql := fmt.Sprintf("select user_id from users %s", con)
	fmt.Println(sql)
	// 查找符合条件的第一行
	row := database.QueryRowDB(sql)
	user_id := -1
	// 在第一行中扫描找到id的字段
	row.Scan(&user_id)
	return user_id
}

// 根据邮箱查询id
func QueryUserWithEmail(email string) int {
	sql := fmt.Sprintf("where email = '%s'", email)
	return QueryUserWightCon(sql)
}

// 根据邮箱和密码来查询id
func QueryUserWithParam(email, password string) int {
	sql := fmt.Sprintf("where email='%s' and password='%s'", email, password)
	return QueryUserWightCon(sql)
}

func Get_user_id(c *gin.Context) int {
	session := sessions.Default(c)
	email := session.Get("loginuser")
	if email == nil {
		return -1
	}
	Email, ok := email.(string)
	if !ok {
		return -1
	}
	user_id := QueryUserWithEmail(Email)
	return user_id
}
