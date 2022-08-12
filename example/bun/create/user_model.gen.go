// Code generated by sqlgen. DO NOT EDIT!

package model

import (
	"context"
	"fmt"
	"time"

	"github.com/uptrace/bun"
)

// UserModel represents a user model.
type UserModel struct {
	db bun.IDB
}

// User represents a user struct data.
type User struct {
	bun.BaseModel `bun:"table:user"`
	Id            uint64    `bun:"id,pk,autoincrement;" json:"id"`
	Name          string    `bun:"name" json:"name"`
	Password      string    `bun:"password" json:"password"`
	Mobile        string    `bun:"mobile" json:"mobile"`
	Gender        string    `bun:"gender" json:"gender"`
	Nickname      string    `bun:"nickname" json:"nickname"`
	Type          int8      `bun:"type" json:"type"`
	CreateTime    time.Time `bun:"create_time" json:"createTime"`
	UpdateTime    time.Time `bun:"update_time" json:"updateTime"`
}

// NewUserModel creates a new user model.
func NewUserModel(db bun.IDB) *UserModel {
	return &UserModel{
		db: db,
	}
}

// Create creates  user data.
func (m *UserModel) Create(ctx context.Context, data ...*User) error {
	if len(data) == 0 {
		return fmt.Errorf("data is empty")
	}

	var list []User
	for _, v := range data {
		list = append(list, *v)
	}

	_, err := m.db.NewInsert().Model(&list).Exec(ctx)
	return err
}