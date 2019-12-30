package controller

import (
	"cake.com/m/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type ClassifyController struct {
	Cxt     iris.Context
	Service service.ClassifyService
}

func (cc *ClassifyController) GetList() mvc.Result {
	list := cc.Service.GetList()
	if len(list) < 0 {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  "1",
				"type":    "",
				"message": "not found",
			},
		}
	}
	return mvc.Response{
		Object: list,
	}
}
