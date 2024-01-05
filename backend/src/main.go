package main

import (
	"backend/src/router"
)

func main() {
	r := router.Router()
	err := r.Run(":8080")
	if err != nil {
		println("gin服务启动失败")
	}
}
