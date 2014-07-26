package helper

import (
	"database/sql"
	"github.com/cloudtrends/revel"
	"github.com/coopernurse/gorp"
	. "golanger.com/middleware"
	"log"
)

type DomoloRevelApp struct {
	*revel.Controller
	//revel.EmptyPlugin
}

func (c DomoloRevelApp) GetDb() *sql.DB {
	log.Println("begin get db obj in middle ware ")
	return Middleware.Get("db").(*sql.DB)
}
func (c DomoloRevelApp) GetDbm() *gorp.DbMap {
	return Middleware.Get("dbm").(*gorp.DbMap)
}
