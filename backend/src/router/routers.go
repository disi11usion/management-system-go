package router

import (
	"backend/src/api"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	api.InitUserAPI(r)
	return r
}
