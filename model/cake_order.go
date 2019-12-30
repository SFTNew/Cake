package model

type CakeOrder struct {
	id          int64   `xorm:"pk autoincr" json:"id"`
	userId      int64   `xorm:"index" json:"user_id"`
	status      int     `xorm:"defualt 0" json:"status"`
	total       float64 `xorm:"float" json:"total"`
	addId       int64   `xorm:"indnex" json:"add_id"`
	serial      string  `xorm:"varchar(25)" json:"serial" `
	remark      string  `xorm:"varchar(100)" json:"remark"`
	payStatus   int     `xorm:"defualt 0" json:"pay_status"`
	deliveryFee float64 `xorm:"float" json:"delivery_fee"`
}
