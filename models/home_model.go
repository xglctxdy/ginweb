package models

import (
	"bytes"
	"html/template"
	"strconv"
)

type HomeBlockParam struct {
	Question_id int
	User_id     int
	Title       string
	Content     string
	CreateTime  string
	// 跳转到该文章的地址
	Link string
}

// 首页显示内容
func MakeHomeBlocks(questions []QuestionWithId) template.HTML {
	htmlHome := ""
	for _, q := range questions {
		//将数据库model转换为首页模板所需要的model
		homeParam := HomeBlockParam{}
		homeParam.Question_id = q.Question_id
		homeParam.User_id = q.User_id
		homeParam.Title = q.Title
		homeParam.Content = q.Content
		homeParam.CreateTime = q.Created_at

		homeParam.Link = "/question/show/" + strconv.Itoa(q.Question_id)

		// 打开模板
		t, _ := template.ParseFiles("views/home_block.html")
		// 存储模板执行后产生的HTML代码
		buffer := bytes.Buffer{}
		// 将模板中的变量替换为homeParam中的数据
		t.Execute(&buffer, homeParam)
		htmlHome += buffer.String()
	}
	return template.HTML(htmlHome)
}
