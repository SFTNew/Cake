package controller

import (
	"cake.com/m/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type OrderDetailsController struct {
	Cxt     iris.Context
	Service service.OrderDetailsService
}

func (odc *OrderDetailsController) GetList() mvc.Result {
	list := odc.Service.GetList()
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
