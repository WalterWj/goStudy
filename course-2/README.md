# 作业
在数据库建一个表 user(id, name, email,age),插入N条记录。
1. 用 go 实现从表里读出数据，并输出到一个文件每行一条记录，用逗号分隔。
2. 要求：
   1. 插入N条数据到表里用一个 goroutine 实现
   2. 读表逻辑用一个 goroutine 来实现
   3. 输出到一个文件的逻辑用 goroutine 来实现；
   4. 插入逻辑完成后通知读表逻辑读取数据
   5. 读表逻辑读出一行记录后通知输出逻辑输出，
   6. 读完后程序全部正确退出。


# 个人学习
go 连接 MySQL 读写
```go
package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "time"
)

//数据库连接信息
const (
    USERNAME = "root"
    PASSWORD = "123qwe"
    NETWORK  = "tcp"
    SERVER   = "127.0.0.1"
    PORT     = 3306
    DATABASE = "test"
)

//user表结构体定义
type User struct {
    Id         int    `json:"id" form:"id"`
    Username   string `json:"username" form:"username"`
    Password   string `json:"password" form:"password"`
    Status     int    `json:"status" form:"status"` // 0 正常状态， 1删除
    Createtime int64  `json:"createtime" form:"createtime"`
}

func CreateTable(DB *sql.DB) {
    sql := `CREATE TABLE IF NOT EXISTS users(
        id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
        username VARCHAR(64),
        password VARCHAR(64),
        status INT(4),
        createtime INT(10)
    );`

    if _, err := DB.Exec(sql); err != nil {
        fmt.Println("create table failed:", err)
        return
    }
    fmt.Println("create table successd")
}

//插入数据
func InsertData(DB *sql.DB) {
    result, err := DB.Exec("insert INTO users(username,password) values(?,?)", "demo", "123qwe")
    if err != nil {
        fmt.Printf("Insert data failed,err:%v", err)
        return
    }
    lastInsertID, err := result.LastInsertId() //获取插入数据的自增ID
    if err != nil {
        fmt.Printf("Get insert id failed,err:%v", err)
        return
    }
    fmt.Println("Insert data id:", lastInsertID)

    rowsaffected, err := result.RowsAffected() //通过RowsAffected获取受影响的行数
    if err != nil {
        fmt.Printf("Get RowsAffected failed,err:%v", err)
        return
    }
    fmt.Println("Affected rows:", rowsaffected)
}

//查询单行
func QueryOne(DB *sql.DB) {
    user := new(User) //用new()函数初始化一个结构体对象
    row := DB.QueryRow("select id,username,password from users where id=?", 1)
    //row.scan中的字段必须是按照数据库存入字段的顺序，否则报错
    if err := row.Scan(&user.Id, &user.Username, &user.Password); err != nil {
        fmt.Printf("scan failed, err:%v\n", err)
        return
    }
    fmt.Println("Single row data:", *user)
}

//查询多行
func QueryMulti(DB *sql.DB) {
    user := new(User)
    rows, err := DB.Query("select id,username,password from users where id = ?", 2)

    defer func() {
        if rows != nil {
            rows.Close() //关闭掉未scan的sql连接
        }
    }()
    if err != nil {
        fmt.Printf("Query failed,err:%v\n", err)
        return
    }
    for rows.Next() {
        err = rows.Scan(&user.Id, &user.Username, &user.Password) //不scan会导致连接不释放
        if err != nil {
            fmt.Printf("Scan failed,err:%v\n", err)
            return
        }
        fmt.Println("scan successd:", *user)
    }
}

//更新数据
func UpdateData(DB *sql.DB) {
    result, err := DB.Exec("UPDATE users set password=? where id=?", "111111", 1)
    if err != nil {
        fmt.Printf("Insert failed,err:%v\n", err)
        return
    }
    fmt.Println("update data successd:", result)

    rowsaffected, err := result.RowsAffected()
    if err != nil {
        fmt.Printf("Get RowsAffected failed,err:%v\n", err)
        return
    }
    fmt.Println("Affected rows:", rowsaffected)
}

//删除数据
func DeleteData(DB *sql.DB) {
    result, err := DB.Exec("delete from users where id=?", 1)
    if err != nil {
        fmt.Printf("Insert failed,err:%v\n", err)
        return
    }
    fmt.Println("delete data successd:", result)

    rowsaffected, err := result.RowsAffected()
    if err != nil {
        fmt.Printf("Get RowsAffected failed,err:%v\n", err)
        return
    }
    fmt.Println("Affected rows:", rowsaffected)
}

func main() {
    conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
    fmt.Println("conn: ", conn)
    DB, err := sql.Open("mysql", conn)
    if err != nil {
        fmt.Println("connection to mysql failed:", err)
        return
    }

    DB.SetConnMaxLifetime(100 * time.Second) //最大连接周期，超时的连接就close
    DB.SetMaxOpenConns(100)                  //设置最大连接数
    CreateTable(DB)
    InsertData(DB)
    QueryOne(DB)
    QueryMulti(DB)
    UpdateData(DB)
    DeleteData(DB)
}
```


# 作业结果

数据库中创建表
```shell
(root@127.0.0.1) [test]>show create table user\G
*************************** 1. row ***************************
       Table: user
Create Table: CREATE TABLE `user` (
  `id` int(11) PRIMARY KEY AUTO_INCREMENT NOT NULL,
  `name` varchar(20) DEFAULT NULL,
  `email` varchar(50) DEFAULT NULL,
  `age` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin
1 row in set (0.00 sec)
```

数据库连接方式
```shell
mysql -ugotest -P 4201 -h127.0.0.1 -p'gotest'
```