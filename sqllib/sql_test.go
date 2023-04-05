package sqllib

import (
	"fmt"
	"testing"
)

/*
	func TestMysqlR(t *testing.T) {
		var testtable IMysql
		testtable = &SMysql{}
		testtable.R().Init("root:dfjns@tcp(127.0.0.1:3306)/wallet?parseTime=true", 600, 1800, 2, 10)
		db := testtable.R().GetDB()
		rows, err := db.Query("SELECT id FROM user")
		if err != nil {
			fmt.Print(err.Error())
			return
		}
		for rows.Next() {
			var Id int
			rows.Scan(&Id)
			fmt.Printf("%v\n", Id)
		}
		rows.Close()
	}

	func TestMysqlW(t *testing.T) {
		var testtable IMysql
		testtable = &SMysql{}
		testtable.WR().Init("root:dfjns@tcp(127.0.0.1:3306)/wallet?parseTime=true", 600, 1800, 2, 10)
		db, err := testtable.WR().GetTX()
		if err != nil {
			fmt.Print(err.Error())
			return
		}
		rows, err := db.Query("SELECT id FROM user")
		if err != nil {
			fmt.Print(err.Error())
			return
		}
		for rows.Next() {
			var Id int
			rows.Scan(&Id)
			fmt.Printf("%v\n", Id)
		}
		rows.Close()
		db.Commit()
	}
*/
func TestPostgresqlW(t *testing.T) {
	testtable := New(SQLType(postgresql))
	testtable.WR().Init("postgresql://root:fgghja@127.0.0.1/test?sslmode=disable", 600, 1800, 2, 10)
	db, err := testtable.WR().GetTX()
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	rows, err := db.Query("select id from test.user")
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	for rows.Next() {
		var Id int
		rows.Scan(&Id)
		fmt.Printf("%v\n", Id)
	}
	rows.Close()
	db.Commit()
}

func TestPostgresqlR(t *testing.T) {
	testtable := New(SQLType(postgresql))
	testtable.R().Init("postgresql://root:fgghja@127.0.0.1/test?sslmode=disable", 600, 1800, 2, 10)
	db := testtable.R().GetDB()
	rows, err := db.Query("select id from test.user")
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	for rows.Next() {
		var Id int
		rows.Scan(&Id)
		fmt.Printf("%v\n", Id)
	}
	rows.Close()
}
