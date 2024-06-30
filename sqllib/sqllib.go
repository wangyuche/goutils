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
	MysqlType      SQLType = "mysql"
	PostgresqlType SQLType = "postgresql"
)

func New(t SQLType) ISQL {
	switch t {
	case MysqlType:
		instace = &SMysql{}
	case PostgresqlType:
		instace = &SPostgresql{}
	default:
		panic("Log Type Error")
	}
	return instace
}
