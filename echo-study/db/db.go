package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	// database driver
	_ "github.com/lib/pq"
)

type database struct {
	conn    *sqlx.DB
	connStr string
}

// DB database pointer var
var DB *database

const (
	host     = "local.ubuntu2"
	port     = "5432"
	user     = "postgres"
	password = "1"
	dbname   = "postgres"
	sslmode  = "disable"
)

// InitDataBase db connect
func InitDataBase() error {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host,
		port,
		user,
		password,
		dbname,
		sslmode)

	DB = &database{connStr: psqlInfo}

	if err := DB.Connect(); nil != err {
		return err
	}

	defer DB.Disconnect()

	return nil
}

func (selfDB *database) Connect() error {
	if nil == selfDB.conn {
		db, err := sqlx.Connect("postgres", selfDB.connStr)
		if nil != err {
			return err
		}
		if err := db.Ping(); nil != err {
			return err
		}

		db.SetMaxOpenConns(200)
		selfDB.conn = db
	}

	return nil
}

func (selfDB *database) Disconnect() error {
	if nil != selfDB.conn {
		if err := selfDB.conn.Close(); nil != err {
			return err
		}

		selfDB.conn = nil
	}

	return nil
}

func (selfDB *database) Get(dest interface{}, query string, args ...interface{}) error {
	if nil == selfDB.conn {
		if err := selfDB.Connect(); nil != err {
			return err
		}
	}

	return selfDB.conn.Get(dest, query, args...)
}

func (selfDB *database) Select(data interface{}, query string, args ...interface{}) error {
	if nil == selfDB.conn {
		if err := selfDB.Connect(); nil != err {
			return err
		}
	}

	return selfDB.conn.Select(data, query, args...)
}
