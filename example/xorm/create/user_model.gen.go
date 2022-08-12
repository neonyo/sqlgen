// Code generated by sqlgen. DO NOT EDIT!

package model

import (
	"context"
	"fmt"
	"time"

	"xorm.io/xorm"
)

// UserModel represents a user model.
type UserModel struct {
	engine xorm.EngineInterface
}

// User represents a user struct data.
type User struct {
	Id         uint64    `xorm:"pk autoincr 'id'" json:"id"`
	Name       string    `xorm:"'name'" json:"name"`
	Password   string    `xorm:"'password'" json:"password"`
	Mobile     string    `xorm:"'mobile'" json:"mobile"`
	Gender     string    `xorm:"'gender'" json:"gender"`
	Nickname   string    `xorm:"'nickname'" json:"nickname"`
	Type       int8      `xorm:"'type'" json:"type"`
	CreateTime time.Time `xorm:"'create_time'" json:"createTime"`
	UpdateTime time.Time `xorm:"'update_time'" json:"updateTime"`
}

func (User) TableName() string {
	return "user"
}

// NewUserModel returns a new user model.
func NewUserModel(engine xorm.EngineInterface) *UserModel {
	return &UserModel{engine: engine}
}

// Insert creates  user data.
func (m *UserModel) Insert(ctx context.Context, data ...*User) error {
	if len(data) == 0 {
		return fmt.Errorf("data is empty")
	}

	var session = m.engine.Context(ctx)
	var list []User
	for _, v := range data {
		list = append(list, *v)
	}

	_, err := session.Insert(&list)
	return err
}