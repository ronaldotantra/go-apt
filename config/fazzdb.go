package config

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/payfazz/go-apt/pkg/fazzdb"
	"sync"
)

var Parameter = fazzdb.Config{
	Limit:  0,
	Offset: 0,
	Lock:   fazzdb.LO_NONE,
}

var DbConf = map[string]string{
	"DB_HOST": "localhost",
	"DB_PORT": "5432",
	"DB_USER": "postgres",
	"DB_PASS": "cashfazz",
	"DB_NAME": "fazz_example",
}

var db *sqlx.DB
var once sync.Once

func GetDB() *sqlx.DB {
	once.Do(func() {
		var err error
		conn := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			DbConf["DB_HOST"],
			DbConf["DB_PORT"],
			DbConf["DB_USER"],
			DbConf["DB_PASS"],
			DbConf["DB_NAME"],
		)
		db, err = sqlx.Connect("postgres", conn)
		if nil != err {
			panic(err)
		}
	})
	return db
}