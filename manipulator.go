package dbcore

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func NewDml() Dml {
	return &MainQuery{}
}

func (q *MainQuery) SelectAll() {
	q.action = "SELECT"
	q.columns = []string{"*"}
}

func (q *MainQuery) SelectColumns(columns []string) {
	q.action = "SELECT"
	q.columns = append(q.columns, columns...)
}

func (q *MainQuery) SelectColumn(column string) {
	q.action = "SELECT"
	q.columns = append(q.columns, column)
}

func (q *MainQuery) Insert() {
	q.action = "INSERT"
}

func (q *MainQuery) Update(tableName string) {
	q.action = "UPDATE"
	q.tableName = tableName
}

func (q *MainQuery) Delete() {
	q.action = "DELETE"
}

func (q *MainQuery) From(tableName string) {
	q.tableName = tableName
}

func (q *MainQuery) Into(tableName string) {
	q.tableName = tableName
}

func (q *MainQuery) Value(column string, value string) {
	q.columns = append(q.columns, column)
	if len(q.insertValues) == 0 {
		q.insertValues = append(q.insertValues, []string{value})
	} else {
		q.insertValues[0] = append(q.insertValues[0], value)
	}
}

func (q *MainQuery) Values(columns []string, values ...string) {
	if len(columns) > 0 {
		q.columns = columns
	}
	q.insertValues = append(q.insertValues, values)
}

func (q *MainQuery) Set(column string, value string) {
	q.conditions = append(q.conditions, condition{"", column, "=", value})
}

func (q *MainQuery) Join(joinType string, joinTables string) {
	q.joinType = append(q.joinType, joinType)
	q.joinTables = append(q.joinTables, joinTables)
}

func (q *MainQuery) On(mainColumn string, joinColumn string, operator string) {
	q.joinCondition = append(q.joinCondition, joinCondition{mainColumn, joinColumn, operator})
}

func (q *MainQuery) Where(joint string, column string, operator string, value string) {
	if checkOperator(operator) {
		q.conditions = append(q.conditions, condition{joint, column, operator, value})
	}
}

func (q *MainQuery) GroupBy(columns []string) {
	q.groupBy = columns
}

func (q *MainQuery) Having(joint string, column string, operator string, value string) {
	if checkOperator(operator) {
		q.having = append(q.having, condition{joint, column, operator, value})
	}
}

func (q *MainQuery) OrderBy(column string, order string) {
	if order == "ASC" || order == "asc" || order == "DESC" || order == "desc" {
		q.orderBy = append(q.orderBy, orderBy{column, order})
	}
}

func (q *MainQuery) Limit(limit int) {
	q.limit = limit
}

func (q MainQuery) Execute(d Database) (result []map[string]string) {
	Log("Run query... : ", q.action)
	query := q.buildQuery()

	if query == "" {
		return
	}

	for !d.IsConnected() {
		time.Sleep(100 * time.Millisecond)
	}
	rows, err := d.db.Query(query)
	if err != nil {
		Error(err)
		return
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		Error(err)
		return
	}

	for rows.Next() {
		columns := make([]string, len(cols))
		columnPointers := make([]interface{}, len(cols))
		data := make(map[string]string)

		for i := range columns {
			columnPointers[i] = &columns[i]
		}

		_ = rows.Scan(columnPointers...)

		for i, colName := range cols {
			data[colName] = columns[i]
		}

		result = append(result, data)
	}

	return result
}

func (q *MainQuery) Clear() {
	q.action = ""
	q.columns = []string{}
	q.tableName = ""
	q.joinType = []string{}
	q.joinTables = []string{}
	q.joinCondition = []joinCondition{}
	q.conditions = []condition{}
	q.groupBy = []string{}
	q.having = []condition{}
	q.orderBy = []orderBy{}
	q.insertValues = [][]string{}
	q.limit = 0
	q.buildQuery()
}

func (q MainQuery) buildQuery() (query string) {
	query = q.action + " "

	if q.action == "SELECT" {
		for i := 0; i < len(q.columns); i++ {
			query += q.columns[i]
			if i != len(q.columns)-1 {
				query += ", "
			}
		}

		if query != "" {
			query += " FROM " + q.tableName
		}

		for i := 0; i < len(q.joinType); i++ {
			query += " " + q.joinType[i] + " " + q.joinTables[i]
			query += " ON " + q.joinCondition[i].mainColumn + " " + q.joinCondition[i].operator + " " + q.joinCondition[i].joinColumn
		}

		query += queryWhere(q)

		if len(q.groupBy) > 0 {
			query += " GROUP BY "
			for i := 0; i < len(q.groupBy); i++ {
				query += q.groupBy[i]
				if i != len(q.groupBy)-1 {
					query += ", "
				}
			}
		}

		if len(q.having) > 0 {
			query += " HAVING "
			for i := 0; i < len(q.having); i++ {
				if i != 0 && i != len(q.having) {
					query += " " + q.having[i].joint + " "
				}

				query += q.having[i].column + " " + q.having[i].operator + " "
				if _, ok := q.having[i].value.(int); ok {
					query += fmt.Sprintf("%v", q.having[i].value)
				} else {
					query += "'" + q.having[i].value.(string) + "'"
				}
			}
		}

		if len(q.orderBy) > 0 {
			query += " ORDER BY "
			for i := 0; i < len(q.orderBy); i++ {
				query += q.orderBy[i].column + " " + q.orderBy[i].order
				if i != len(q.orderBy)-1 {
					query += ", "
				}
			}
		}

		if q.limit > 0 {
			query += " LIMIT " + fmt.Sprintf("%v", q.limit)
		}
	} else if q.action == "INSERT" {
		if len(q.columns) != len(q.insertValues[0]) {
			Error(errors.New(ERROR_INVALID_QUERY))
		}

		query += "INTO " + q.tableName

		if len(q.insertValues) > 0 {
			for i := 0; i < len(q.columns); i++ {
				if i == 0 {
					query += " ("
				}

				query += q.columns[i]

				if i != len(q.columns)-1 {
					query += ", "
				} else {
					query += ") VALUES "
				}
			}

			for i, insertValue := range q.insertValues {
				query += "("

				for j, value := range insertValue {
					query += "'" + value + "'"
					if j < len(insertValue)-1 {
						query += ", "
					} else {
						query += ")"
					}
				}

				if i < len(q.insertValues)-1 {
					query += ", "
				}
			}
		}
	} else if q.action == "UPDATE" {
		query += q.tableName
		query += " SET " + q.conditions[0].column + " = '" + q.conditions[0].value.(string) + "'"
		query += queryWhere(q)
	} else if q.action == "DELETE" {
		query += " FROM " + q.tableName
		query += " WHERE " + q.conditions[0].column + " = '" + q.conditions[0].value.(string) + "'"
	}

	if !isValidQuery(query) {
		query = ""
		Error(errors.New(ERROR_INVALID_QUERY))
	}

	return
}

func (q MainQuery) GetQueryString() string {
	return q.buildQuery()
}

func checkOperator(operator string) bool {
	switch operator {
	case "=", "!=", ">", "<", ">=", "<=", "LIKE", "NOT LIKE", "IN", "NOT IN", "BETWEEN", "NOT BETWEEN":
		return true
	}
	return false
}

func isValidQuery(q string) bool {
	return strings.Count(q, ";") == 0
}

func queryWhere(q MainQuery) (query string) {
	if len(q.conditions) > 0 {
		bracketOpen := true
		query = " WHERE ("
		nextJoint := ""
		start := 0

		if q.action == "UPDATE" {
			start = 1
		}

		for i := start; i < len(q.conditions); i++ {
			if q.action == "UPDATE" && i == 0 {
				continue
			}

			if i != len(q.conditions)-1 {
				nextJoint = q.conditions[i+1].joint
			}

			if i != start && i != len(q.conditions) {
				query += " " + q.conditions[i].joint + " "
			}

			if (q.conditions[i].joint == "AND" && nextJoint == "OR") && i != start {
				query += "("
				bracketOpen = true
			}

			query += q.conditions[i].column + " " + q.conditions[i].operator + " "

			switch q.conditions[i].operator {
			case "IN", "NOT IN":
				query += "('" + q.conditions[i].value.(string) + "')"
			default:
				query += "'" + q.conditions[i].value.(string) + "'"
			}

			if q.conditions[i].joint != "AND" && nextJoint == "AND" {
				query += ")"
				bracketOpen = false
			}
		}
		if bracketOpen {
			query += ")"
		}
	}
	return query
}
