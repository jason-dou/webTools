package main

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"time"
	_ "webTools/models"
	_ "webTools/routers"
)

func init() {
	dbUser, _ := web.AppConfig.String("dbuser")
	dbPassword, _ := web.AppConfig.String("dbpassword")
	dbName, _ := web.AppConfig.String("dbname")
	dataSource := fmt.Sprintf("%v:%v@/%v", dbUser, dbPassword, dbName)

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dataSource)

	// Set to UTC time
	orm.DefaultTimeLoc = time.UTC
}

func main() {
	web.Run()
}

