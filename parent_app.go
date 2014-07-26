package helper

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	//. "golanger.com/middleware"
)

type ParentApp interface {
	GetDb() *sql.DB
	GetDbm() *gorp.DbMap
}
