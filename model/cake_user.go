package model

import "time"

type CakeUser struct {
	id         int64
	name       string
	password   string
	header     string
	sex        int
	token      string
	integral   float64
	sessionKey string
	openId     string
	createTime time.Time
	updateTime time.Time
}
