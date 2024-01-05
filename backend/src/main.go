package main

import (
	"backend/src/router"
)

func main() {
	r := router.Router()
	_ = r.SetTrustedProxies([]string{"127.0.0.1"})
	err := r.Run(":8080")
	if err != nil {
		println("gin服务启动失败")
	}
}
