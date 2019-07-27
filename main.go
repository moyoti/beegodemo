package main

import (
	_ "demo/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:my110120130@tcp(127.0.0.1:3306)/test")
}
func main() {
	beego.Run()
}
