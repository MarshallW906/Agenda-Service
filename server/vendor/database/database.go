package database

import (
	"args"
	"entity"
	er "err"

	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"

	// register the sqlite3 engine for go
	_ "github.com/mattn/go-sqlite3"
)

// Engine of xorm
var Engine *xorm.Engine

func init() {
	// initialize xorm engine
	engine, err := xorm.NewEngine("sqlite3", *args.DBfile)
	er.CheckErr(err)
	engine.SetMapper(core.SameMapper{})
	err = engine.Sync2(new(entity.User), new(entity.Token))
	er.CheckErr(err)
	Engine = engine
}
