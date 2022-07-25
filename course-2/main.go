package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//数据库连接信息
const (
	USERNAME = "gotest"
	PASSWORD = "gotest"
	NETWORK  = "tcp"
	SERVER   = "127.0.0.1"
	PORT     = 4201
	DATABASE = "test"
)

// 定义 user 表的结构体
// type User struct {
// 	id    int    `json:"id", from: 1`
// 	name  string `json:"name", from: name`
// 	email string `json:"email", from: xxx@xxx.com`
// 	age   int    `json:"age", from: 18`
// }

//插入数据
func InsertData(db *sql.DB, data []map[string]string) {
	// result, err := DB.Exec("insert INTO user(name,email,age) values(?,?,?)", "demo", "123qwe", 1)
	sqlStr := "INSERT INTO user(name, email, age) VALUES "
	vals := []interface{}{}
	for _, row := range data {
		sqlStr += "(?, ?, ?),"
		vals = append(vals, row["name"], row["email"], row["age"])
	}
	// 去掉结尾 , 号，SQL 最后拼接 ; 号
	sqlStr = sqlStr[:len(sqlStr)-1] + ";"
	// 数据库 prepare SQL
	insert, _ := db.Prepare(sqlStr)
	defer insert.Close()
	// 开启事务
	begin, _ := db.Begin()
	// 拼接 SQL：insert into values(),(),()
	result, err := begin.Stmt(insert).Exec(vals...)
	// 异常处理和输出插入数据
	if err != nil {
		fmt.Printf("Insert data failed,err:%v", err)
		return
	}
	// 事务提交
	err = begin.Commit()
	if err != nil {
		fmt.Println("事务提交失败:", err)
		return
	}
	// 通过RowsAffected获取受影响的行数
	rowsaffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v", err)
		return
	}
	fmt.Println("插入数据成功，Affected rows:", rowsaffected)
}

func main() {
	// 初始化数据
	var data = []map[string]string{
		{"name": "一号", "email": "1@qq.com", "age": "18"},
		{"name": "二号", "email": "2@qq.com", "age": "18"},
		{"name": "三号", "email": "3@qq.com", "age": "18"},
	}
	// 连接数据库
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		return
	}
	// 插入数据库
	InsertData(db, data)
}
