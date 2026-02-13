package main

import (
	"fmt"
	"system/dao"
	"system/router"
)

func main() {
	dao.InitMySQL()
	fmt.Println("数据库连接成功")

	r := router.SetUpRouter()

	fmt.Println("xx，启动！")
	r.Run()
}
