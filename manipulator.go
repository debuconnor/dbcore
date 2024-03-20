package dbcore

import (
	"fmt"
	"strings"
)

func checkOperator(operator string) bool {
	switch operator {
	case "=", "!=", ">", "<", ">=", "<=", "LIKE", "NOT LIKE", "IN", "NOT IN", "BETWEEN", "NOT BETWEEN":
		return true
	}
	return false
}

func isValidQuery(q string) bool {
	return strings.Count(q, ";") <= 1
}

func (q *MainQuery) Clear() {
	q.columns = []string{}
	q.tableName = ""
	q.joinType = []string{}
	q.joinTables = []string{}
	q.joinCondition = []joinCondition{}
	q.conditions = []condition{}
	q.groupBy = []string{}
	q.having = []condition{}
	q.orderBy = []orderBy{}
	q.buildQuery()
}

func (q *MainQuery) SelectColumns(columns []string) {
	q.columns = append(q.columns, columns...)
}

func (q *MainQuery) SelectColumn(column string) {
	q.columns = append(q.columns, column)
}

func (q *MainQuery) SelectAll() {
	q.columns = []string{"*"}
}

func (q *MainQuery) From(tableName string) {
	q.tableName = tableName
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

func (q *MainQuery) buildQuery() (query string) {
	for i := 0; i < len(q.columns); i++ {
		if i == 0 {
			query = "SELECT "
		}

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

	if len(q.conditions) > 0 {
		query += " WHERE "

		for i := 0; i < len(q.conditions); i++ {
			if i != 0 && i != len(q.conditions) {
				query += " " + q.conditions[i].joint + " "
			}

			query += q.conditions[i].column + " " + q.conditions[i].operator + " "

			switch q.conditions[i].operator {
			case "IN", "NOT IN":
				query += "('" + q.conditions[i].value.(string) + "')"
			default:
				query += "'" + q.conditions[i].value.(string) + "'"
			}
		}
	}

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

	if !isValidQuery(query) {
		query = ""
	}

	return
}

func (q *MainQuery) Execute(d Database) (result []map[string]string) {
	rows, _ := d.db.Query(q.buildQuery())
	fmt.Println(q.buildQuery())
	defer rows.Close()

	cols, _ := rows.Columns()

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
