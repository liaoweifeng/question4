package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8")
	db.SetMaxOpenConns(1000)
	err := db.Ping()
	if err != nil {
		fmt.Println("fail to connect to db")
	}
}

func DBConn() *sql.DB {
	return db
}

//通过用户名和密码完成user表中注册操作
func UserSignup(username string, password string) bool {
	sqlStr := `select username from user where username=?;`
	rowobj := db.QueryRow(sqlStr, username)
	var name string
	rowobj.Scan(&name)
	if username != name {
		stmt, err := DBConn().Prepare(
			"insert into user(username,password,signday,power)values(?,?,0,0) ")
		if err != nil {
			fmt.Println("fail to insert")
			return true
		}
		defer stmt.Close()

		_, err = stmt.Exec(username, password)
		if err != nil {
			fmt.Println("fail to insert")
			return true
		}
	} else {
		return true
	}
	return false
}

//判断用户登录账号密码是否正确
func UserSignin(username string, password string)bool {
	sqlStr := "select username,password from user where username=?;"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return false
	}
	defer stmt.Close()
	rows, err := stmt.Query(username)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return false
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		var passwd string
		err := rows.Scan(&name, &passwd)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return false
		}
		if name == username && passwd == password {
			return true
		} else{return false}
	}
	return false
}
