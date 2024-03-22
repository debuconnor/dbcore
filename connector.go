package dbcore

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewDb() Connection {
	return &Database{}
}

func (d *Database) SetConnection(host, port, username, password, dbName string) {
	d.host = host
	d.port = port
	d.username = username
	d.password = password
	d.dbName = dbName
}

func (d *Database) ConnectMysql() {
	Log("Connecting to MySQL...")
	var err error
	d.db, err = sql.Open("mysql", d.username+":"+d.password+"@tcp("+d.host+":"+d.port+")/"+d.dbName)

	if err != nil {
		Log("Error while connecting to MySQL.")
	} else {
		Log("Connected to MySQL.")
	}
}

func (d *Database) DisconnectMysql() {
	d.db.Close()
	Log("Disconnected from MySQL.")
}

func (d *Database) IsConnected() bool {
	return d.db.Ping() == nil
}

func (d Database) GetDb() Database {
	return d
}
