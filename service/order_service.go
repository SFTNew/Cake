package service

import (
	"cake.com/m/model"
	"github.com/go-xorm/xorm"
)

type OrderService interface {
	GetList() []model.CakeOrder
}

type orderService struct {
	Engine *xorm.Engine
}

func NewOrderService(engine *xorm.Engine) OrderService {
	return &orderService{
		Engine: engine,
	}
}

func (os *orderService) GetList() []model.CakeOrder {
	var list = make([]model.CakeOrder, 0)
	return list
}
