package model

import "time"

type CakeClassify struct {
	id         int64     `xorm:"pk autoincr" json:"id"`
	name       string    `xorm:"varchar(100)" json:"name"`
	level      int       `xorm:"default 0" json:"level"`
	status     int       `xorm:"default 0" json:"status"`
	createTime time.Time `xorm:"DateTime" json:"create_time"`
	updateTime time.Time `xorm:"DateTime" json:"update_time"`
}
