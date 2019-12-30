package service

import (
	"cake.com/m/model"
	"github.com/go-xorm/xorm"
)

type ConfigService interface {
	GetList() []model.CakeConfig
}

type configService struct {
	Engine *xorm.Engine
}

func NewConfigService(engine *xorm.Engine) ConfigService {
	return &configService{
		Engine: engine,
	}
}

func (cs *configService) GetList() []model.CakeConfig {
	list := make([]model.CakeConfig, 0)
	cs.Engine.Where("status=0").Find(&list)
	return list
}
