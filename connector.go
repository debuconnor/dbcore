package dbcore

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func (d *Database) SetConnection(host, port, username, password, dbName string) {
	d.host = host
	d.port = port
	d.username = username
	d.password = password
	d.dbName = dbName
}

func (d *Database) ConnectMysql() error {
	var err error
	d.db, err = sql.Open("mysql", d.username+":"+d.password+"@tcp("+d.host+":"+d.port+")/"+d.dbName)
	return err
}

func (d *Database) DisconnectMysql() {
	d.db.Close()
}

func (d Database) IsConnected() bool {
	return d.db.Ping() == nil
}
