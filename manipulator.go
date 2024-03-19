package dbcore

func (q *mainQuery) SelectColumns(columns []string) {
	q.columns = columns
}

func (q *mainQuery) From(tableName string) {
	q.tableName = tableName
}

func (q *mainQuery) Join(joinType string, joinTables string) {
	q.joinType = append(q.joinType, joinType)
	q.joinTables = append(q.joinTables, joinTables)
}

func (q *mainQuery) On(mainColumn string, joinColumn string, operator string) {
	q.joinCondition = append(q.joinCondition, joinCondition{mainColumn, joinColumn, operator})
}

func (q *mainQuery) Where(conditions []condition) {
	q.conditions = conditions
}

func (q *mainQuery) GroupBy(columns []string) {
	q.groupBy.columns = columns
}

func (q *mainQuery) Having(conditions []condition) {
	q.having = conditions
}

func (q *mainQuery) OrderBy(columns []string, order string) {
	if order != "ASC" && order != "asc" && order != "DESC" && order != "desc" {
		q.orderBy = append(q.orderBy, orderBy{columns, order})
	}
}

func checkOperator(operator string) bool {
	if operator != "=" && operator != "!=" && operator != ">" && operator != "<" && operator != ">=" && operator != "<=" {
		return false
	}
	return true
}

func (q *mainQuery) buildQuery() (query string) {
	query = "SELECT "
	for i := 0; i < len(q.columns); i++ {
		query += q.columns[i]
		if i != len(q.columns)-1 {
			query += ", "
		}
	}

	query += " FROM " + q.tableName

	for i := 0; i < len(q.joinType); i++ {
		query += " " + q.joinType[i] + " " + q.joinTables[i]
		query += " ON " + q.joinCondition[i].mainColumn + " " + q.joinCondition[i].operator + " " + q.joinCondition[i].joinColumn
	}

	if len(q.conditions) > 0 {
		query += " WHERE "
		for i := 0; i < len(q.conditions); i++ {
			query += q.conditions[i].column + " " + q.conditions[i].operator + " " + q.conditions[i].value.(string)
			if i != len(q.conditions)-1 {
				query += " AND "
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
			query += q.having[i].column + " " + q.having[i].operator + " " + q.having[i].value.(string)
			if i != len(q.having)-1 {
				query += " AND "
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
	return
}
