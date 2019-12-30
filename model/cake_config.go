package model

type CakeConfig struct {
	id     int64  `xorm:"pk autoincr" json:"id"`
	cKey   string `xorm:"varchar(20)" json:"c_key"`
	cValue string `xorm:"varchar(20)" json:"c_value"`
	info   string `xorm:"varchar(100)" json:"info"`
	status int    `xorm:"default 0" json:"status"`
}
