package service

import (
	"cake.com/m/model"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
)

type BannerService interface {
	GetList() []model.CakeBanner
}

func NewBannerService(engine *xorm.Engine) BannerService {
	return &bannerService{
		Engine: engine,
	}
}

type bannerService struct {
	Engine *xorm.Engine
}

func (bs *bannerService) GetList() []model.CakeBanner {
	list := make([]model.CakeBanner, 0)
	err := bs.Engine.Where("status = 0").Find(&list)
	if err != nil {
		iris.New().Logger().Info(err.Error())
	}
	return list
}
