package model

import "time"

type CakeOrderDetails struct {
	id         int64
	goodId     int64
	orderId    int64
	name       string
	price      float64
	number     int64
	totalPrice float64
	status     int
	createTime time.Time
	updateTime time.Time
}
