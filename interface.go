package dbcore

type Dml interface {
	SelectAll()
	SelectColumns([]string)
	SelectColumn(string)
	Insert()
	Update(string)
	Delete()
	From(string)
	Into(string)
	Value(string, string)
	Set(string, string)
	Join(string, string)
	On(string, string, string)
	Where(string, string, string, string)
	GroupBy([]string)
	Having(string, string, string, string)
	OrderBy(string, string)
	Execute(Database) []map[string]string
	Clear()
	buildQuery() string
	GetQueryString() string
}

type Ddl interface {
	CreateTable() error
	AlterTable() error
	DropTable() error
}

type Connection interface {
	SetConnection(string, string, string, string, string)
	ConnectMysql() error
	DisconnectMysql()
	IsConnected() bool
	GetDb() Database
}
