package db

import (
	"github.com/azd1997/go-frame/config"
	_ "github.com/go-sql-driver/mysql"		// 使用mysql必须
	"github.com/go-xorm/xorm"
	"github.com/pkg/errors"
)

var engine *xorm.Engine

func InitEngine(conf config.DBConfig) error {
	var err error

	engine, err = xorm.NewEngine("mysql",
		conf.DbUser + ":" + conf.DbPassword +
		"@tcp(" + conf.DbHost + ":" + conf.DbPort + ")/" +
		conf.DbName + "?charset=utf8")
	if err != nil {
		return errors.Wrap(err, "InitEngine: NewEngine failed: ")
	}

	err = engine.Ping()
	if err != nil {
		return errors.Wrap(err, "InitEngine: PingEngine failed: ")
	}

	return nil
}

func GetEngine() *xorm.Engine {
	return engine
}
