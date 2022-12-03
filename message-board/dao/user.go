package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	dns := "root:zx1913683154@tcp(127.0.0.1:3306)/user"
	DB, _ = sql.Open("mysql", dns)
	err := DB.Ping()
	if err != nil {
		fmt.Printf("connect failed,err:%v\n", err)
		return
	}
	fmt.Println("connect successfully")
}
