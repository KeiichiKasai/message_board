package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"message-board/api/middleware"
	"message-board/dao"
	"message-board/model"
	"net/http"
)

func register(c *gin.Context) {
	username := c.PostForm("username") //注册用户名
	password := c.PostForm("password") //注册密码
	telephone := c.PostForm("telephone")
	rows, err := dao.DB.Query("select username from user where id>?", 0)
	if err != nil {
		fmt.Printf("Query failed,err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u model.User
		err := rows.Scan(&u.Username)
		if err != nil {
			fmt.Printf("Scan failed,err:%v\n", err)
			return
		}
		if username == u.Username {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "用户已存在",
			})
			return
		}
	}
	_, err = dao.DB.Exec("insert into user(username,password,telephone)values(?,?,?)", username, password, telephone)
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "注册成功",
	})
}

func login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	rows, err := dao.DB.Query("select username,password from user where id>?", 0)
	if err != nil {
		fmt.Scanf("Query failed,err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u model.User
		err = rows.Scan(&u.Username, &u.Password)
		if err != nil {
			fmt.Printf("Scan failed,err:%v\n", err)
			return
		}
		if username == u.Username && password == u.Password {
			tokenString, _ := middleware.GenToken(u.Username)
			c.JSON(http.StatusOK, gin.H{
				"message": "登录成功",
				"data":    gin.H{"token": tokenString},
			})
			return
		}
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "用户名或密码错误",
	})
}

func forget(c *gin.Context) {
	username := c.PostForm("username")
	newpassword := c.PostForm("newpassword")
	telephone := c.PostForm("telephone")
	rows, err := dao.DB.Query("select username,telephone from user where id>?", 0)
	if err != nil {
		fmt.Printf("Query failed,err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u model.User
		err = rows.Scan(&u.Username, &u.Telephone)
		if err != nil {
			fmt.Printf("Scan failed,err:%v\n", err)
			return
		}
		if username == u.Username && telephone == u.Telephone {
			_, err := dao.DB.Exec("update user set password=? where username=?", newpassword, u.Username)
			if err != nil {
				fmt.Printf("Update failed,err:%v\n", err)
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"message": "修改密码成功",
			})
			return
		}

	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "用户名与手机号不匹配，修改失败",
	})
}
