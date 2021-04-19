package models

type TplLevel struct {
	Id  int `xorm:"not null pk autoincr comment('当前的等级') INT(11)"`
	Exp int `xorm:"default 0 comment('升级需要的经验') INT(11)"`
}
