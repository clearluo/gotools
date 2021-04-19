package models

import (
	"time"
)

type TUserGame struct {
	Id       int       `xorm:"not null pk autoincr INT(11)"`
	UserId   int       `xorm:"not null default 0 comment('用户Id') INT(11)"`
	GameId   int       `xorm:"not null default 0 comment('游戏Id') INT(11)"`
	Count    int       `xorm:"not null default 0 INT(11)"`
	CreateAt time.Time `xorm:"default CURRENT_TIMESTAMP TIMESTAMP"`
}
