package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var DB *sql.DB

func InitDB() {

	dns := fmt.Sprintf("%s:%s@tcp(%s%s)/%s",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.database"),
	)
	DB, _ = sql.Open("mysql", dns)
	err := DB.Ping()
	if err != nil {
		fmt.Printf("connect failed,err:%v\n", err)
		return
	}
	fmt.Println("connect successfully")
}
