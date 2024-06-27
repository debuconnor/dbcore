package dbcore

import "database/sql"

type Database struct {
	db       *sql.DB
	host     string
	port     string
	username string
	password string
	dbName   string
	platform string
}

type MainQuery struct {
	action        string
	columns       []string
	tableName     string
	joinType      []string
	joinTables    []string
	joinCondition []joinCondition
	conditions    []condition
	groupBy       []string
	having        []condition
	orderBy       []orderBy
	insertValues  [][]string
	limit         int
}

type joinCondition struct {
	mainColumn string
	joinColumn string
	operator   string
}

type condition struct {
	joint    string
	column   string
	operator string
	value    interface{}
}

type orderBy struct {
	column string
	order  string
}

type Schema struct {
	tableAction     string
	tableName       string
	columnAction    string
	columns         []string
	revisedColumns  []string
	varType         []string
	isNull          []bool
	isPk            []bool
	isAutoIncrement []bool
	defaultVal      []string
	comment         []string
}
