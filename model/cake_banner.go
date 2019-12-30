package model

import "time"

type CakeBanner struct {
	id         int64     `xorm:"pk autoincr" josn:"id"`
	img        string    `xorm:"varchar(100)" json:"img"`
	name       string    `xorm:"varchar(100)" json:"name"`
	url        string    `xorm:"varchar(100)" json:"url"`
	status     int64     `xorm:"default 0" json:"status"`
	createTime time.Time `xorm:"DateTime" json:"create_time"`
}
