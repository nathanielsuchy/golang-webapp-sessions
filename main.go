package main

import (
	"fmt"
	"webapp/models"
	_ "webapp/routers"

	"github.com/beego/beego/v2/adapter/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// Create Database Connection
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root@/orm_test?charset=utf8")

	// Register Models
	orm.RegisterModel(new(models.User))

	// Sync Database to ORM Models
	name := "default" // database alias

	// Drop table and re-create.
	force := false

	// Print log.
	verbose := true

	// Error.
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	beego.Run()
}
