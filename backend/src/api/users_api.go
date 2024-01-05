package api

import (
	"backend/src/model"
	"backend/src/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitUserAPI(g *gin.Engine) {
	c := &UsersApi{}
	c.init(g) //这里需要引入你的gin框架的实例
	fmt.Println("user服务启动成功")
}

func (t UsersApi) init(g *gin.Engine) {
	// 依次: 分页列表，单条，插入，修改，删除
	group := g.Group("/users")
	group.GET("/list", t.list) //不设置限制条件的画默认查询所有
	group.GET("/one", t.one)
	group.POST("/insert", t.insert)
	group.POST("/update", t.update)
	group.POST("/delete", t.delete)
}

// UsersApi 控制器
type UsersApi struct {
	Page int   `form:"page,default=1"`
	Size int   `form:"size,default=10"`
	Ids  []int `form:"ids"`
}

// 分页列表
func (t UsersApi) list(c *gin.Context) {
	_ = c.Bind(&t)
	v := model.Users{}
	_ = c.Bind(&v)
	c.JSON(http.StatusOK, model.OkData(service.UsersService.List(t.Page, t.Size, &v)))
}

// 根据主键Id查询记录
func (t UsersApi) one(c *gin.Context) {
	var v model.Users
	_ = c.Bind(&v)
	fmt.Printf("%d\n", v.Id)
	c.JSON(http.StatusOK, model.OkData(service.UsersService.One(v.Id)))
}

// 修改记录
func (t UsersApi) update(c *gin.Context) {
	var v model.Users
	_ = c.ShouldBindJSON(&v)
	service.UsersService.Update(v)
	c.JSON(http.StatusOK, model.OkMsg("修改成功！"))
}

// 插入记录
func (t UsersApi) insert(c *gin.Context) {
	var v model.Users
	_ = c.ShouldBindJSON(&v)
	service.UsersService.Insert(v)
	c.JSON(http.StatusOK, model.OkMsg("插入成功！"))
}

// 根据主键删除
func (t UsersApi) delete(c *gin.Context) {
	_ = c.Bind(&t)
	service.UsersService.Delete(t.Ids)
	c.JSON(http.StatusOK, model.OkMsg("删除成功！"))
}
