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

// TestStruct role table
type TestStruct struct {
	RoleID   string `db:"role_id" json:"role_id"`
	RoleName string `db:"role_name" json:"role_name"`
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

	DB := &database{connStr: psqlInfo}

	if err := DB.Connect(); nil != err {
		return err
	}

	data := TestStruct{}
	err := DB.Get(&data, "SELECT role_id, role_name FROM role")

	if nil != err {
		fmt.Println(err)
	}

	fmt.Println(data)

	return nil
}

func (selfDB *database) Connect() error {
	fmt.Println(selfDB.conn)

	if selfDB.conn == nil {

		db, err := sqlx.Connect("postgres", selfDB.connStr)
		if nil != err {
			return err
		}
		if err := db.Ping(); nil != err {
			return err
		}

		db.SetMaxOpenConns(200)
		fmt.Println(db)
		selfDB.conn = db
		fmt.Println(selfDB.conn)
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

func (selfDB *database) Get(data interface{}, query string, args ...interface{}) error {
	if err := selfDB.Connect(); nil != err {
		return err
	}

	return selfDB.conn.Get(data, query, args...)
}
