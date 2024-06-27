package dbcore

import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
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

func (d *Database) SetConnectionFromGcpSecret(secretVersion string) {
	resultJson := accessSecretVersion(secretVersion)
	dbInfo := parseJson(resultJson)

	d.SetConnection(dbInfo["host"].(string), dbInfo["port"].(string), dbInfo["username"].(string), dbInfo["password"].(string), dbInfo["dbname"].(string))
}

func (d *Database) ConnectMysql() error {
	Log("Connecting to MySQL...")
	var err error
	d.db, err = sql.Open("mysql", d.username+":"+d.password+"@tcp("+d.host+":"+d.port+")/"+d.dbName)

	if err != nil {
		Log("Error while connecting to MySQL.")
		return err
	} else {
		Log("Connected to MySQL.")
		d.platform = "mysql"
	}

	return nil
}

func (d *Database) ConnectMssql() error {
	Log("Connecting to MSSQL...")
	var err error
	d.db, err = sql.Open("mssql", "server="+d.host+";user id="+d.username+";password="+d.password+";port="+d.port+";database="+d.dbName+";encrypt=disable")

	if err != nil {
		Log("Error while connecting to MSSQL.")
		return err
	} else {
		Log("Connected to MSSQL.")
		d.platform = "mssql"
	}

	return nil
}

func (d *Database) Disconnect() {
	d.db.Close()
	Log("Disconnected from Database Server.")
}

func (d *Database) IsConnected() bool {
	return d.db.Ping() == nil
}

func (d Database) GetDb() Database {
	return d
}
