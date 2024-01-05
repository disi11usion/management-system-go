package service

import (
	"backend/src/common"
	"backend/src/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var UsersService = usersService{}

// usersService 业务层
type usersService struct {
}

func (t usersService) db() *gorm.DB {

	return common.GetDB()
}

// List 分页列表
func (t usersService) List(page, size int, v *model.Users) gin.H {
	lists := make([]model.Users, 0) // 结果
	t.db().Model(&model.Users{}).Where(&v).Order("").Offset((page - 1) * size).Limit(size).Find(&lists)
	var total int64 // 统计
	t.db().Model(&v).Where(&v).Count(&total)
	h := gin.H{"list": lists, "total": total}
	return h
}

// One 根据主键查询
func (t usersService) One(id interface{}) (v model.Users) {
	t.db().Find(&v, id)
	return
}

// Update 修改记录
func (t usersService) Update(v model.Users) {
	t.db().Model(&v).Updates(v)
}

// Insert 插入记录
func (t usersService) Insert(v model.Users) {
	t.db().Save(&v)
}

// Delete 根据主键删除
func (t usersService) Delete(ids []int) {
	t.db().Delete(model.Users{}, ids)
}
