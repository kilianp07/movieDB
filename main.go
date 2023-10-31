package main

import (
	"fmt"
	"os"

	_ "github.com/kilianp07/movieDB/routers"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	sqlconn, err := beego.AppConfig.String("sqlconn")
	if err != nil {
		fmt.Println("Cannot find sqlconn in config file:", err)
		os.Exit(1)
	}

	// Initialisez l'ORM avec l'alias "default" en utilisant la valeur sqlconn
	orm.RegisterDataBase("default", "mysql", sqlconn)
	orm.RunSyncdb("default", false, true)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
