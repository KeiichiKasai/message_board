package main

import (
	"message-board/api"
	"message-board/api/middleware"
	"message-board/dao"
)

func main() {
	middleware.ViperSetup()
	middleware.InitZap()
	dao.InitDB()
	api.InitRouter()
}
