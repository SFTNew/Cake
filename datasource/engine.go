package dataSource

import (
	_ "github.com/go-sql-driver/mysql" //不能忘记导入
	"github.com/go-xorm/xorm"
)

func NewMySqlEngine() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:root@/cake?charset=utf8")

	// err = engine.Sync2(new(
	// 	model.Permission),
	// 	new(model.City),
	// 	new(model.Admin),
	// 	new(model.AdminPermission),
	// 	new(model.User),
	// 	new(model.UserOrder))
	if err != nil {
		panic(err.Error())
	}

	//设置显示Sql语句
	engine.ShowSQL(true)
	engine.SetMaxOpenConns(10)
	return engine
}
