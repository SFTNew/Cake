package service

import (
	"cake.com/m/model"
	"github.com/go-xorm/xorm"
)

type OrderDetailsService interface {
	GetList() []model.CakeOrderDetails
}

type orderDetailsService struct {
	Engine *xorm.Engine
}

func NewOrderDetailsService(engine *xorm.Engine) OrderDetailsService {
	return &orderDetailsService{
		Engine: engine,
	}
}

func (oc *orderDetailsService) GetList() []model.CakeOrderDetails {
	var list = make([]model.CakeOrderDetails, 0)
	return list
}
