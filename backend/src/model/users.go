package model

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Users struct {
	//
	Id int64 `gorm:"primaryKey" json:"id,omitempty" form:"id"`
	//
	UserName string `json:"user_name,omitempty" form:"user_name"`
	//
	AvatarUrl string `json:"avatar_url,omitempty" form:"avatar_url"`
	//
	Gender int64 `json:"gender,omitempty" form:"gender"`
	//
	Password string `json:"password" form:"password"`
	//
	Phone string `json:"phone,omitempty" form:"phone"`
	//
	Email string `json:"email,omitempty" form:"email"`
	//
	CreatedAt time.Time `json:"-" form:"created_at"`
	//
	UpdatedAt time.Time `json:"-" form:"updated_at"`
	//
	UserStatus int64 `json:"user_status,omitempty" form:"user_status"`
	//
	IfDelete int64 `json:"if_delete,omitempty" form:"if_delete"`
}

func OkData[T any](t T) any {
	return t
}

func OkMsg(s string) any {
	return gin.H{"status": s}
}

func ErrMsg(s string) any {
	return gin.H{"error": s}

}
