package models

type TplItem struct {
	Id         int    `xorm:"not null pk autoincr INT(11)"`
	Name       string `xorm:"default '' comment('道具的名字') VARCHAR(255)"`
	GroupId    int    `xorm:"default 0 comment('组的id') INT(11)"`
	MinScore   int    `xorm:"default 0 comment('获得的最低分数') INT(11)"`
	MaxScore   int    `xorm:"default 0 comment('获得的最高分数') INT(11)"`
	RebornTime int    `xorm:"default 0 comment('重生的时间') INT(11)"`
	RebornMinR int    `xorm:"default 0 comment('重生的最小半径') INT(11)"`
	RebornMaxR int    `xorm:"default 0 comment('重生的最大半径') INT(11)"`
}
