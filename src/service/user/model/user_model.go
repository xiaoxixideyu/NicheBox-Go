package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Uid            int64  `gorm:"index:idx_user_uid"`
	Email          string `gorm:"index:idx_user_email_pwd,priority:1;index:uniq_user_email,unique"`
	Password       string `gorm:"index:idx_user_email_pwd,priority:2"`
	Username       string
	Introduction   string
	AvatarOriginId int64
	AvatarWebpId   int64
	Telephone      string
	ShowCollection bool
	ShowFollower   bool
	ShowFollwing   bool
	ShowBox        bool
	ShowContent    bool
}
