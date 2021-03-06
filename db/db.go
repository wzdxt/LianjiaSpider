package db

import (
	_ "github.com/go-sql-driver/mysql"
	"sync"
	"database/sql"
)

var instance *sql.DB
var once sync.Once

func DBInstance() *sql.DB {
	once.Do(func() {
		var err error
		if instance, err = sql.Open("mysql", "root:123456@(lianjia-mysql:3306)/lianjia?parseTime=true"); err != nil {
			panic(err)
		}
	})
	return instance
}
