package models

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

func init() {
	dbName := beego.AppConfig.String("dbName")

	engine, err := xorm.NewEngine("mysql", dbName)
	if err != nil {
		beego.Critical(`Open Database error`)
		os.Exit(-1)
	}

	engine.ShowDebug = true
	engine.ShowErr = true
	engine.ShowSQL = true
	engine.ShowWarn = true

	f, err := os.OpenFile("sql.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		beego.Warn(`Create sql log error`)
	} else {
		engine.Logger = xorm.NewSimpleLogger(f)
	}

	engine.Sync2(new(Leak), new(Patch))
	Engine = engine
}

var Engine *xorm.Engine
