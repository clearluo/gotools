package models

type TplGame struct {
	Id          int    `xorm:"not null pk autoincr INT(11)"`
	Name        string `xorm:"default '' comment('标题') VARCHAR(50)"`
	Icon        string `xorm:"default '' comment('列表展示图片') VARCHAR(255)"`
	BgImg       string `xorm:"default '' comment('背景图片') VARCHAR(255)"`
	Level       int    `xorm:"default 0 comment('困难度') INT(11)"`
	ItemGroupId int    `xorm:"default 0 comment('物品组id') INT(11)"`
	Duration    int    `xorm:"default 300 comment('游戏的时长') INT(11)"`
	WinExp      int    `xorm:"default 0 comment('胜利经验') INT(11)"`
	LossExp     int    `xorm:"default 0 comment('失败经验') INT(11)"`
	WinHonor    int    `xorm:"default 0 comment('获胜的荣誉') INT(11)"`
	LossHonor   int    `xorm:"default 0 comment('失败的荣誉') INT(11)"`
}
