package model

type CakeAdd struct {
	id      int64  `xorm:"pk autoincr" json:"id"`
	userId  int64  `xorm:"index" json:"user_id"`
	address string `xorm:"varchar(100)" json:"address"`
	mobile  string `xorm:"varchar(11)" json:"mobile"`
	status  int64  `xorm:"default 0" json:"status"`
	name    string `xorm:"varchar(18)" json:"name"`
	city    string `xorm:"varchar(100)" json:"city"`
}
