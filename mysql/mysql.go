package mysql

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type IWR interface {
	Init(string, int, int, int, int)
	GetDB() *sql.DB
	GetTX() (*sql.Tx, error)
}

type IR interface {
	Init(string, int, int, int, int)
	GetDB() *sql.DB
}

type IMysql interface {
	WR() IWR
	R() IR
}

type SMysql struct {
	SWR *SWR
	SR  *SR
}

func (this *SMysql) WR() IWR {
	if this.SWR == nil {
		this.SWR = &SWR{}
	}
	return this.SWR
}

func (this *SMysql) R() IR {
	if this.SR == nil {
		this.SR = &SR{}
	}
	return this.SR
}

type SR struct {
	DB *sql.DB
}

func (this *SR) Init(parameter string, MaxIdleTimeSecond, MaxLifetimeSecond, MaxIdleConns, MaxOpenConns int) {
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

func (this *SR) GetDB() *sql.DB {
	return this.DB
}

type SWR struct {
	DB *sql.DB
}

func (this *SWR) Init(parameter string, MaxIdleTimeSecond, MaxLifetimeSecond, MaxIdleConns, MaxOpenConns int) {
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

func (this *SWR) GetDB() *sql.DB {
	return this.DB
}

func (this *SWR) GetTX() (*sql.Tx, error) {
	return this.DB.Begin()
}
