package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Uid            int64  `gorm:"index:idx_user_uid"`
	Email          string `gorm:"index:idx_user_email_pwd,priority:1;index:idx_user_email,unique"`
	Password       string `gorm:"index:idx_user_email_pwd,priority:2"`
	Username       string
	Avatar         string
	Telephone      string
	ShowCollection bool
	ShowFollower   bool
	ShowFollwing   bool
	ShowBox        bool
	ShowContent    bool
}
