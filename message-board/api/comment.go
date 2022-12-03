package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/dao"

)

func comment(c *gin.Context) {
	message_id := c.PostForm("message_id")
	nickname := c.PostForm("nickname")
    comme:=c.PostForm("comment")
	c.PostForm("comment")
	_, err := dao.DB.Exec("insert into comment (message_id,nickname,comment)values (?,?,?)",message_id,nickname,comme)
	if err!=nil{
		fmt.Printf("insert failed,err:%v\n",err)
		c.JSON(500,gin.H{
			"message":"评论失败",
		})
	}
	c.JSON(200,gin.H{
		"message":"评论成功",
	})
}
