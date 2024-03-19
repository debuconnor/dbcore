package dbcore

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func (d *database) setConnection(host, port, username, password, dbName string) {
	d.host = host
	d.port = port
	d.username = username
	d.password = password
	d.dbName = dbName
}

func (d *database) connectMysql() error {
	var err error
	d.db, err = sql.Open("mysql", d.username+":"+d.password+"@tcp("+d.host+":"+d.port+")/"+d.dbName)
	return err
}

func (d *database) disconnectMysql() error {
	return d.db.Close()
}

func (d *database) isConnected() bool {
	return d.db.Ping() == nil
}
