package sqllib

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type SPostgresql struct {
	SWR *WRpostgresql
	SR  *Rpostgresql
}

func (this *SPostgresql) WR() IWR {
	if this.SWR == nil {
		this.SWR = &WRpostgresql{}
	}
	return this.SWR
}

func (this *SPostgresql) R() IR {
	if this.SR == nil {
		this.SR = &Rpostgresql{}
	}
	return this.SR
}

type Rpostgresql struct {
	DB *sql.DB
}

func (this *Rpostgresql) Init(parameter string, MaxIdleTimeSecond, MaxLifetimeSecond, MaxIdleConns, MaxOpenConns int) {
	db, err := sql.Open("postgres", parameter)
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

func (this *Rpostgresql) GetDB() *sql.DB {
	return this.DB
}

type WRpostgresql struct {
	DB *sql.DB
}

func (this *WRpostgresql) Init(parameter string, MaxIdleTimeSecond, MaxLifetimeSecond, MaxIdleConns, MaxOpenConns int) {
	db, err := sql.Open("postgres", parameter)
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

func (this *WRpostgresql) GetDB() *sql.DB {
	return this.DB
}

func (this *WRpostgresql) GetTX() (*sql.Tx, error) {
	return this.DB.Begin()
}
