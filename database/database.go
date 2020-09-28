package database

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//DBSession Global for DB conn
var DBSession *sql.DB

var dbSchema = "XIssueTracker"

func init() {
	var err error
	//short sleep waiting for DB
	time.Sleep(time.Duration(13) * time.Second)

	mysqlhost := os.Getenv("MYSQL_HOST")
	if mysqlhost == "" {
		mysqlhost = "192.168.2.90"
	}

	mysqlusername := os.Getenv("MYSQL_USER")
	if mysqlusername == "" {
		mysqlusername = "root"
	}

	mysqlpassword := os.Getenv("MYSQL_PASSWORD")
	if mysqlpassword == "" {
		mysqlpassword = "chuck111"
	}

	DBSession, err = sql.Open("mysql", mysqlusername+":"+mysqlpassword+"@tcp("+mysqlhost+":3306)/"+dbSchema+"?parseTime=true")

	if err != nil {
		log.Print("DB CONN ERROR!")
		panic(err.Error())
	}

	if err = DBSession.Ping(); err != nil {
		log.Print("DB PING ERROR!")
		panic(err.Error())
	}
}
