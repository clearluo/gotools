package models

import (
	"time"
)

type TUser struct {
	Id           int       `xorm:"not null pk autoincr INT(11)"`
	Username     string    `xorm:"default '' comment('用户名') unique VARCHAR(64)"`
	Password     string    `xorm:"default '' comment('密码') VARCHAR(64)"`
	Level        int       `xorm:"default 0 comment('等级') INT(11)"`
	Exp          int       `xorm:"default 0 comment('经验') INT(11)"`
	Honor        int       `xorm:"default 0 comment('荣誉') INT(11)"`
	InstanceId   int       `xorm:"default 0 comment('当前副本id') INT(11)"`
	LastTickTime int       `xorm:"default 0 comment('最后同步时间') INT(11)"`
	SkinId       int       `xorm:"default 0 INT(11)"`
	CreateAt     time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME"`
	UpdateAt     time.Time `xorm:"default CURRENT_TIMESTAMP TIMESTAMP"`
}
