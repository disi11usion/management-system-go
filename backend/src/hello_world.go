package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	UserName string `json:"user_name,omitempty"`
	Age      string `json:"age,omitempty"`
}

func Index(ctx *gin.Context) {
	var user = UserInfo{
		UserName: "ysc",
		Age:      "24",
	}
	ctx.JSON(http.StatusOK, user)
}

func main() {

	router := gin.Default()
	router.GET("/index", Index)
	router.Run(":8080")
}
