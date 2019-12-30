package controller

import (
	"cake.com/m/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type OrderController struct {
	Cxt     iris.Context
	Service service.OrderService
}

func (oc *OrderController) GetList() mvc.Result {
	list := oc.Service.GetList()
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
