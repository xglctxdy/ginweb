package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// 一个指向SQL数据库的连接的指针类型
var db *sql.DB

func InitMysql() {
	//初始化数据库
	fmt.Println("InitMysql....")
	if db == nil {
		db, _ = sql.Open("mysql", "root:114514@tcp(192.168.10.133:3306)/gindemo?charset=utf8")
		CreateTableWithUser()
		CreateTableWithQuestions()
		CreateTableWithAnswers()
		CreateTableWithAlbum()
	}
	fmt.Println("Init Done")
}

// 查询并且期望结果只有一行
func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}

// 查询并期望查询结果有多行
func QueryDB(sql string) (*sql.Rows, error) {
	return db.Query(sql)
}

// 操作数据库
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

// 创建用户表
func CreateTableWithUser() {
	sql := `CREATE TABLE IF NOT EXISTS users(
		user_id INT AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(50) NOT NULL,
	    password VARCHAR(100) NOT NULL,
   		email VARCHAR(100) UNIQUE NOT NULL,
   		avatar VARCHAR(255),
   		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   		account_status INT DEFAULT 0
		);`
	ModifyDB(sql)
}

// 创建问题表
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

// 创建回答表
func CreateTableWithAnswers() {
	sql := `create table if not exists answers(
		answer_id INT AUTO_INCREMENT PRIMARY KEY,
    	question_id INT NOT NULL,
    	user_id INT NOT NULL,
    	content TEXT NOT NULL,
    	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    	FOREIGN KEY (question_id) REFERENCES questions(question_id),
    	FOREIGN KEY (user_id) REFERENCES users(user_id)
    );`
	ModifyDB(sql)
}

// 创建图片表
func CreateTableWithAlbum() {
	sql := `create table if not exists album(
		image_id INT AUTO_INCREMENT PRIMARY KEY,
		related_id INT NOT NULL,
		related_type ENUM('question', 'answer') NOT NULL,
		image_path VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (related_id) REFERENCES questions(question_id) ON DELETE CASCADE,
		FOREIGN KEY (related_id) REFERENCES answers(answer_id) ON DELETE CASCADE
		);`
	ModifyDB(sql)
}
