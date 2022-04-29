package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB
var ID int64

func main() {
	db, _ := initDB()
	error := db.Ping()
	if error != nil {
		fmt.Println("数据库连接失败!")
	} else {
		fmt.Println("数据库连接成功!")
	}
	//Insert(db)
	Update(db)
	//Delete(db)
	//Query(db)
}

func initDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./test.db")
	//db, err := sql.Open("xxx", "mygosqlte.db")
	if err != nil {
		fmt.Println("数据库初始化失败!")
		return nil, err
	}
	fmt.Println("数据库初始化成功!")
	//fmt.Print(db)
	return db, err
}

// 增
func Insert(db *sql.DB) {
	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) VALUES(?,?,?)")
	checkErr(err)

	res, err := stmt.Exec("超超", "研发部门", "2012-12-09")
	fmt.Println("插入成功", res)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)
	ID = id

	fmt.Println(id)
}

// 删
func Delete(db *sql.DB) {
	stmt, err := db.Prepare("DELETE FROM userinfo WHERE uid=?")
	checkErr(err)

	res, err := stmt.Exec(1)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()
	fmt.Println("删除成功", res)
}

// 改
func Update(db *sql.DB) {
	stmt, err := db.Prepare("UPDATE userinfo SET username=? WHERE uid=?")
	checkErr(err)

	res, err := stmt.Exec("星爷", 2)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
	fmt.Println("修改成功", res)
}

// 查
func Query(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
