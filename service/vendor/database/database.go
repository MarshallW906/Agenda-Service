package database

import (
	"args"
	"entity"
	er "err"

	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"

	_ "github.com/mattn/go-sqlite3"
)

// Engine of xorm
var Engine *xorm.Engine

func init() {
	// initialize xorm engine
	Engine, err := xorm.NewEngine("sqlite3", *args.DBfile)
	er.CheckErr(err)
	Engine.SetMapper(core.SameMapper{})
	err = Engine.Sync2(new(entity.User))
	er.CheckErr(err)
}
