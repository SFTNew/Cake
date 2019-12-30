package model

import "time"

type CakeGood struct {
	id            int64     `xorm:"pk autoincr" json:"id"`
	name          string    `xorm:"varchar(200)" json:"name"`
	profile       string    `xorm:"varchar(255)" json:"profile"`
	specification string    `xorm:"varchar(50)" json:"specification"`
	stock         int64     `xorm:"defualt 10000" json:"stock"`
	price         float64   `xorm:"float" json:"price"`
	img           string    `xorm:"varchar(100)" json:"img"`
	classifyId    int64     `xorm:"index" json:"img"`
	status        int       `xorm:"defualt 0" json:"status"`
	createTime    time.Time `xorm:"DateTime" json:"create_time"`
	updateTime    time.Time `xorm:"DateTime" json:"update_time"`
}
