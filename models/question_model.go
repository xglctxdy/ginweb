package models

import (
	"fmt"
	"gindemo/config"
	"gindemo/database"
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

type Question struct {
	User_id int
	Title   string
	Content string
}

type QuestionWithId struct {
	Question_id int
	User_id     int
	Title       string
	Content     string
	Created_at  string
}

func AddQuestion(question Question) (int64, error) {
	i, err := InsertArticle(question)
	return i, err
}

func InsertArticle(question Question) (int64, error) {
	return database.ModifyDB("insert into questions(user_id,title,content) values(?,?,?)",
		question.User_id, question.Title, question.Content)
}

// 根据页码查询文章
func FindQuestionWithPage(page int) ([]QuestionWithId, error) {
	page--
	fmt.Println("---------->page", page)
	//从配置文件中获取每页的文章数量
	return QueryQuestionWithPage(page, config.Num)
}

/*
*
分页查询数据库
limit分页查询语句，

	语法：limit m，n
	m代表从多少位开始获取，与id值无关
	n代表获取多少条数据

注意limit前面没有where
*/

func QueryQuestionWithPage(page, num int) ([]QuestionWithId, error) {
	sql := fmt.Sprintf("limit %d,%d", page*num, num)
	return QueryQuestionWithCon(sql)
}

func QueryQuestionWithCon(sql string) ([]QuestionWithId, error) {
	sql = "select question_id,user_id,title,content,created_at from questions " + sql
	rows, err := database.QueryDB(sql)
	if err != nil {
		return nil, err
	}
	var qList []QuestionWithId
	for rows.Next() {
		question_id := 0
		user_id := 0
		title := ""
		content := ""
		create_at := ""
		rows.Scan(&question_id, &user_id, &title, &content, &create_at)
		q := QuestionWithId{question_id, user_id, title, content, create_at}
		qList = append(qList, q)
	}
	return qList, nil
}
