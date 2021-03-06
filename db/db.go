package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB gorm.DB

func init() {
	var err error

	DB, err = gorm.Open("mysql", "savet5:savet5@tcp(database:3306)/savet5?charset=utf8&parseTime=True")

	//fmt.Println(DB)

	if err != nil {
		panic(fmt.Sprintf("Got error when connect database, the error is '%v'", err))
	}
	// Get database connection handle [*sql.DB](http://golang.org/pkg/database/sql/#DB)
	//d := db.DB()

	// With it you could use package `database/sql`'s builtin methods
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	DB.DB().Ping()

	// By default, table name is plural of struct type, you can use struct type as table name with:
	// IvanA: I use plural which is default, but set it anyway just to be sure (if default is ever changed)
	DB.SingularTable(false)

	// we need logging, at least in development, and probbaly later as well:
	DB.LogMode(true)
}
