package service

import (
	"cake.com/m/model"
	"github.com/go-xorm/xorm"
)

type GoodService interface {
	GetList() []model.CakeGood
}

type goodService struct {
	Engine *xorm.Engine
}

func NewGoodService(engine *xorm.Engine) GoodService {
	return &goodService{
		Engine: engine,
	}
}

func (gs *goodService) GetList() []model.CakeGood {
	var list = make([]model.CakeGood, 0)
	return list
}
