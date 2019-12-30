package controller

import (
	"cake.com/m/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type AddressController struct {
	Ctx     iris.Context
	Service service.AddressService
}

func (ac *AddressController) GetList() mvc.Result {
	// v1/add/1
	// path := ac.Ctx.Path()
	// fmt.Println(path)
	// var pathSlice []string
	// if path != "" {
	// 	pathSlice = strings.Split(path, "/")
	// }
	// id, _ := strconv.ParseInt(pathSlice[3], 10, 64)
	addrs := ac.Service.GetList(2)
	if len(addrs) == 0 {
		return mvc.Response{
			Object: map[string]interface{}{
				"status": 1,
				"msg":    "无数据",
				"data":   "",
			},
		}
	}

	return mvc.Response{
		Object: map[string]interface{}{
			"status": 0,
			"msg":    "",
			"data":   &addrs,
		},
	}
}
