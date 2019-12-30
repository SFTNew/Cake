package service

import "github.com/go-xorm/xorm"

import "cake.com/m/model"

import "github.com/kataras/iris"

type ClassifyService interface {
	GetList() []model.CakeClassify
}

type classifyService struct {
	Engine *xorm.Engine
}

func NewCalssifyService(engine *xorm.Engine) ClassifyService {
	return &classifyService{
		Engine: engine,
	}
}

func (cs *classifyService) GetList() []model.CakeClassify {
	list := make([]model.CakeClassify, 0)

	err := cs.Engine.Where("status = 0").Find(&list)

	if err != nil {
		iris.New().Logger().Error(err.Error())
	}

	return list

}
