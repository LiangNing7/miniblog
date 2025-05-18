package model

import (
	"github.com/LiangNing7/goutils/pkg/rid"
	"gorm.io/gorm"
)

// AfterCreate 在创建数据库记录之后生成 postID.
func (m *PostM) AfterCreate(tx *gorm.DB) error {
	PostID := rid.NewResourceID("post")
	m.PostID = PostID.New(uint64(m.ID))

	return tx.Save(m).Error
}

// AfterCreate 在创建数据库记录之后生成 userID.
func (m *UserM) AfterCreate(tx *gorm.DB) error {
	UserID := rid.NewResourceID("user")
	m.UserID = UserID.New(uint64(m.ID))

	return tx.Save(m).Error
}
