package dbcore

import "database/sql"

type database struct {
	db       *sql.DB
	host     string
	port     string
	username string
	password string
	dbName   string
}

type mainQuery struct {
	columns       []string
	tableName     string
	joinType      []string
	joinTables    []string
	joinCondition []joinCondition
	conditions    []condition
	groupBy       []string
	having        []condition
	orderBy       []orderBy
}

type joinCondition struct {
	mainColumn string
	joinColumn string
	operator   string
}

type condition struct {
	column   string
	operator string
	value    interface{}
}

type orderBy struct {
	column string
	order  string
}

type queryResult struct {
	queryString  string
	rowsAffected int64
	rowCount     int64
}
