package dbcore

type dml interface {
	SelectColumns([]string)
	From(string)
	Join(string, string)
	On(string, string, string)
	Where(string, string, string)
	GroupBy([]string)
	Having(string, string, string)
	OrderBy([]string, string)
	buildQuery() string
	Execute(Database) []map[string]interface{}
}

type ddl interface {
	CreateTable() error
	AlterTable() error
	DropTable() error
}

type connection interface {
	setConnection(host, port, username, password, dbName string)
	connectMysql() error
	disconnectMysql() error
	isConnected() bool
}
