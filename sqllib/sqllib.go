package sqllib

import (
	"database/sql"
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

type ISQL interface {
	WR() IWR
	R() IR
}

var instace ISQL

type SQLType string

const (
	mysql      SQLType = "mysql"
	postgresql SQLType = "postgresql"
)

func New(t SQLType) ISQL {
	switch t {
	case mysql:
		instace = &SMysql{}
	case postgresql:
		instace = &SPostgresql{}
	default:
		panic("Log Type Error")
	}
	return instace
}
