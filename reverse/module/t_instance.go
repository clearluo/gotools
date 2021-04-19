package models

import (
	"time"
)

type TInstance struct {
	Id          int       `xorm:"not null pk autoincr INT(11)"`
	GameId      int       `xorm:"comment('游戏的模板id') INT(11)"`
	IsTeam      int       `xorm:"default 0 comment('是否组队') TINYINT(255)"`
	UserId      int       `xorm:"default 0 comment('创建的用户id') INT(11)"`
	TargetId    int       `xorm:"default 0 comment('参与者id') INT(11)"`
	IsWin       int       `xorm:"default 0 comment('是否获胜') TINYINT(255)"`
	UserScore   int       `xorm:"default 0 comment('创建放分数') INT(11)"`
	TargetScore int       `xorm:"default 0 comment('参与者分数') INT(11)"`
	Data        string    `xorm:"TEXT"`
	CreateAt    time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME"`
}
