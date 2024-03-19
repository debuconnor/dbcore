package dbcore

type dml interface {
	SelectColumns([]string)
	From(string)
	Join(string, string)
	On(string, string, string)
	Where([]condition)
	GroupBy([]string)
	Having([]condition)
	OrderBy([]string, string)
	buildQuery() string
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
