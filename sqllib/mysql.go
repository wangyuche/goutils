package sqllib

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type SMysql struct {
	SWR *WRmysql
	SR  *Rmysql
}

func (this *SMysql) WR() IWR {
	if this.SWR == nil {
		this.SWR = &WRmysql{}
	}
	return this.SWR
}

func (this *SMysql) R() IR {
	if this.SR == nil {
		this.SR = &Rmysql{}
	}
	return this.SR
}

type Rmysql struct {
	DB *sql.DB
}

func (this *Rmysql) Init(parameter string, MaxIdleTimeSecond, MaxLifetimeSecond, MaxIdleConns, MaxOpenConns int) {
	db, err := sql.Open("mysql", parameter)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	db.SetConnMaxIdleTime(time.Duration(MaxIdleTimeSecond) * time.Second)
	db.SetConnMaxLifetime(time.Duration(MaxLifetimeSecond) * time.Second)
	db.SetMaxIdleConns(MaxIdleConns)
	db.SetMaxOpenConns(MaxOpenConns)
	this.DB = db
}

func (this *Rmysql) GetDB() *sql.DB {
	return this.DB
}

type WRmysql struct {
	DB *sql.DB
}

func (this *WRmysql) Init(parameter string, MaxIdleTimeSecond, MaxLifetimeSecond, MaxIdleConns, MaxOpenConns int) {
	db, err := sql.Open("mysql", parameter)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	db.SetConnMaxIdleTime(time.Duration(MaxIdleTimeSecond) * time.Second)
	db.SetConnMaxLifetime(time.Duration(MaxLifetimeSecond) * time.Second)
	db.SetMaxIdleConns(MaxIdleConns)
	db.SetMaxOpenConns(MaxOpenConns)
	this.DB = db
}

func (this *WRmysql) GetDB() *sql.DB {
	return this.DB
}

func (this *WRmysql) GetTX() (*sql.Tx, error) {
	return this.DB.Begin()
}
