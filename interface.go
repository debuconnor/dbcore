package dbcore

type Dml interface {
	SelectAll()
	SelectColumns([]string)
	SelectColumn(string)
	SelectFunction(string, ...string)
	Insert()
	Update(string)
	Delete()
	From(string)
	Into(string)
	Value(string, string)
	Values([]string, ...string)
	Set(string, string)
	Join(string, string)
	On(string, string, string)
	Where(string, string, string, ...string)
	GroupBy([]string)
	Having(string, string, string, string)
	OrderBy(string, string)
	Limit(int)
	Execute(Database) []map[string]string
	Clear()
	buildQuery() string
	GetQueryString() string
}

type Ddl interface {
	CheckTableExists(Database, string) bool
	CreateTable(string)
	AlterTable(string)
	DropTable(string)
	AddColumn(string, string, bool, bool, bool, string, string)
	DropColumn(string)
	ChangeColumn(string, string, string, string, string)
	SetColumnDefault(string, string)
	Execute(Database)
	Clear()
	buildQuery() string
	GetQueryString() string
}

type Connection interface {
	SetConnection(string, string, string, string, string)
	SetConnectionFromGcpSecret(string)
	ConnectMysql() error
	ConnectMssql() error
	Disconnect()
	IsConnected() bool
	GetDb() Database
}
