package controllers

import (
	"fmt"
	"gindemo/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
表的定义
func CreateTableWithQuestions() {
	sql := `CREATE TABLE IF NOT EXISTS questions(
		question_id INT AUTO_INCREMENT PRIMARY KEY,
    	user_id INT NOT NULL,
    	title VARCHAR(255) NOT NULL,
    	content TEXT NOT NULL,
    	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    	FOREIGN KEY (user_id) REFERENCES users(user_id)
		);`
	ModifyDB(sql)
}
*/

func AddQuestionGet(c *gin.Context) {
	//获取session
	islogin := GetSession(c)
	// Id 为 0
	c.HTML(http.StatusOK, "add_question.html", gin.H{"IsLogin": islogin})
}

func AddQuestionPost(c *gin.Context) {
	//获取浏览器传输的数据，通过表单的name属性获取值
	//获取表单信息
	title := c.PostForm("title")
	content := c.PostForm("content")
	user_id := models.Get_user_id(c)

	if user_id == -1 {
		fmt.Print("找不到当前登录的用户")
	}

	q := models.Question{user_id, title, content}
	_, err := models.AddQuestion(q)

	response := gin.H{}
	if err == nil {
		//无误
		response = gin.H{"code": 0, "message": "ok"}
	} else {
		response = gin.H{"code": 1, "message": "error"}
	}

	c.JSON(http.StatusOK, response)

}
