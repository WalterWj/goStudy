package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"

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
type User struct {
	id    int
	name  string
	email string
	age   int
}

//插入数据
func InsertData(db *sql.DB, data []map[string]string, c chan bool) {
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
	_, err := begin.Stmt(insert).Exec(vals...)
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

	// 以上全部成功，输出给 c 为 true
	c <- true
}

//查询多行
func QueryMulti(DB *sql.DB, c chan string) {
	user := new(User)
	rows, err := DB.Query("select id,name,email,age from user;")
	if err != nil {
		fmt.Printf("Query failed,err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&user.id, &user.name, &user.email, &user.age) // 不scan会导致连接不释放
		if err != nil {
			fmt.Printf("Scan failed,err:%v\n", err)
			return
		}
		fmt.Println("数据为:", *user)
		// content := strconv.Itoa(user.id) + user.name + user.email + strconv.Itoa(user.age)
		content := strings.Join([]string{strconv.Itoa(user.id), user.name, user.email, strconv.Itoa(user.age)}, " ")
		c <- content
	}
	close(c)

}

// 写文件
func WriteFle(content chan string, c chan bool) {
	//创建一个新文件
	filePath := "user.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()
	//循环接收通道的数据
	for {
		ch2strings, ok := <-content
		if !ok {
			//全部写入完成，通知结束
			fmt.Println("写入成功")
			break
		} else {
			//写入文件时，使用带缓存的 *Writer
			//Flush将缓存的文件真正写入到文件中
			write := bufio.NewWriter(file)
			write.WriteString(ch2strings + "\n")
			write.Flush()
		}
	}
	c <- true
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
	defer db.Close()
	// 插入数据库
	c := make(chan bool)
	c1 := make(chan string)
	go InsertData(db, data, c)
	rlt := <-c
	// 如果插入成功，rlt 为 true，进入下一步
	if rlt {
		fmt.Println("数据插入成功，进行数据查询和数据写入文件")
		go QueryMulti(db, c1)
		go WriteFle(c1, c)

		<-c
		close(c)
	} else {
		fmt.Println("数据插入失败")
	}
}
