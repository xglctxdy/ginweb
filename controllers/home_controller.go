package controllers

import (
	"fmt"
	"gindemo/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomeGet(c *gin.Context) {
	// 获取session，检查是否登录
	islogin := GetSession(c)
	page := 1
	var qList []models.QuestionWithId
	qList, err := models.FindQuestionWithPage(page)
	if err != nil {
		fmt.Println(err)
	}
	html := models.MakeHomeBlocks(qList)
	c.HTML(http.StatusOK, "home.html", gin.H{"IsLogin": islogin, "Content": html})
}
