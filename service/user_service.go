package service

import (
	"cake.com/m/model"
	"github.com/go-xorm/xorm"
)

type UserService interface {
	GetList() []model.CakeUser
}

type userService struct {
	Engine *xorm.Engine
}

func NewUserService(engine *xorm.Engine) UserService {
	return &userService{
		Engine: engine,
	}
}

func (uc *userService) GetList() []model.CakeUser {
	var list = make([]model.CakeUser, 0)
	return list
}
