package provider

import (
	"github.com/cocotyty/summer"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DBProvider struct {
	DB      *sqlx.DB
	Address string `sm:"#db.url"`
}

func (d *DBProvider) Init() {
	d.DB = sqlx.MustOpen("mysql", d.Address)
}

func (d *DBProvider) Provider() interface{} {
	return d.DB
}

func init() {
	summer.Add("db", &DBProvider{})
}
