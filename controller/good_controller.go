package controller

import (
	"cake.com/m/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type GoodController struct {
	Cxt     iris.Context
	Service service.GoodService
}

func (gc *GoodController) GetList() mvc.Result {
	list := gc.Service.GetList()
	if len(list) < 0 {
		return mvc.Response{
			Object: map[string]interface{}{
				"status": "",
				"msg":    "",
			},
		}
	}

	return mvc.Response{
		Object: list,
	}
}
