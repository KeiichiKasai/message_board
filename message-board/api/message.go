package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/dao"
	"message-board/model"
	"net/http"
)

func WriteMessage(c *gin.Context) {
	nickname := c.PostForm("nickname")
	message := c.PostForm("message")
	_, err := dao.DB.Exec("insert into message(nikename,message)values (?,?)", nickname, message)
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "写入时发生错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "留言成功",
	})
	return
}
func ResearchMessage(c *gin.Context) {
	var m model.MessageBoard
	id := c.PostForm("id")
	err := dao.DB.QueryRow("select nickname,message,`time` from message where id=?", id).Scan(&m.NickName, &m.Message, &m.Time)
	if err != nil {
		fmt.Printf("select failed,err:%v\n", err)
		c.JSON(500, gin.H{
			"message": "获取留言失败",
		})
	}
	c.JSON(200, gin.H{
		"time":     m.Time,
		"nickname": m.NickName,
		"message":  m.Message,
	})
}
