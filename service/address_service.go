package service

import (
	"cake.com/m/model"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
)

type AddressService interface {
	GetList(id int64) []*model.CakeAdd
}

type addressService struct {
	engine *xorm.Engine
}

func NewAddressService(db *xorm.Engine) AddressService {
	return &addressService{
		engine: db,
	}
}

func (ad *addressService) GetList(id int64) []*model.CakeAdd {
	var addrs []*model.CakeAdd
	err := ad.engine.Where("user_id=?", id).Find(&addrs)
	if err != nil {
		iris.New().Logger().Error(err.Error())
		panic(err.Error())
		return nil
	}
	return addrs
}
