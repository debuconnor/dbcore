package dbcore

func NewDdl() Ddl {
	return &Schema{}
}

func (s *Schema) CheckTableExists(d Database, tableName string) bool {
	SaveLog("", "Checking if schema exists...")
	s.tableAction = "SHOW TABLES LIKE "
	s.tableName = tableName
	query := s.tableAction + "'" + s.tableName + "'"

	rows, _ := d.db.Query(query)
	defer rows.Close()

	for rows.Next() {
		SaveLog("", "Schema found.")
		return true
	}

	SaveLog("", "Schema not found.")
	return false
}

func (s *Schema) CreateTable(tableName string) {
	s.tableAction = "CREATE TABLE"
	s.tableName = tableName
}

func (s *Schema) AlterTable(tableName string) {
	s.tableAction = "ALTER TABLE"
	s.tableName = tableName
}

func (s *Schema) DropTable(tableName string) {
	s.tableAction = "DROP TABLE"
	s.tableName = tableName
}

func (s *Schema) AddColumn(columnName string, varType string, isNull bool, isPk bool, isAutoIncrement bool, defaultVal string, comment string) {
	s.columns = append(s.columns, columnName)
	s.varType = append(s.varType, varType)
	s.isNull = append(s.isNull, isNull)
	s.isPk = append(s.isPk, isPk)
	s.isAutoIncrement = append(s.isAutoIncrement, isAutoIncrement)
	s.defaultVal = append(s.defaultVal, defaultVal)
	s.comment = append(s.comment, comment)
}

func (s *Schema) DropColumn(columnName string) {
	s.columnAction = "DROP COLUMN"
	s.tableName = columnName
}

func (s *Schema) ChangeColumn(column string, changeName string, varType string, defaultVal string, comment string) {
	s.columnAction = "CHANGE COLUMN"
	s.columns = append(s.columns, column)
	s.revisedColumns = append(s.revisedColumns, changeName)
	s.varType = append(s.varType, varType)
	s.defaultVal = append(s.defaultVal, defaultVal)
	s.comment = append(s.comment, comment)
}

func (s *Schema) SetColumnDefault(columnName string, defaultVal string) {
	s.columnAction = "ALTER COLUMN"
	s.columns = append(s.columns, columnName)
	s.defaultVal = append(s.defaultVal, defaultVal)
}

func (s Schema) Execute(d Database) {
	SaveLog("", "Run query... :", s.tableAction)
	query := s.buildQuery()
	_, err := d.db.Exec(query)

	if err != nil {
		SaveLog("", "Error while executing query.")
	}
}

func (s *Schema) Clear() {
	s.tableAction = ""
	s.tableName = ""
	s.columnAction = ""
	s.columns = []string{}
	s.revisedColumns = []string{}
	s.varType = []string{}
	s.isNull = []bool{}
	s.isPk = []bool{}
	s.isAutoIncrement = []bool{}
	s.comment = []string{}
	s.buildQuery()
}

func (s Schema) buildQuery() (query string) {
	SaveLog("", "Building query...")
	switch s.tableAction {
	case "CREATE TABLE":
		query = "CREATE TABLE " + s.tableName + " ("
		for i := 0; i < len(s.columns); i++ {
			query += s.columns[i] + " " + s.varType[i]
			if s.isNull[i] {
				query += " NULL"
			} else {
				query += " NOT NULL"
			}
			if s.isAutoIncrement[i] {
				query += " AUTO_INCREMENT"
			}
			if s.defaultVal[i] != "" {
				query += " DEFAULT '" + s.defaultVal[i] + "'"
			}
			if s.comment[i] != "" {
				query += " COMMENT '" + s.comment[i] + "'"
			}
			query += ", "
		}
		query += "PRIMARY KEY ("
		for i := 0; i < len(s.columns); i++ {
			if s.isPk[i] {
				query += s.columns[i] + ", "
			}
		}
		query = query[:len(query)-2]
		query += "))"

	case "ALTER TABLE":
		query = "ALTER TABLE " + s.tableName + " "
		switch s.columnAction {
		case "DROP COLUMN":
			query += "DROP COLUMN " + s.tableName
		case "CHANGE COLUMN":
			query += "CHANGE COLUMN " + s.columns[0] + " " + s.revisedColumns[0] + " " + s.varType[0]
			if s.comment[0] != "" {
				query += " COMMENT '" + s.comment[0] + "'"
			}
		case "ALTER COLUMN":
			query += "ALTER COLUMN " + s.columns[0] + " SET DEFAULT '" + s.defaultVal[0] + "'"
		}

	case "DROP TABLE":
		query = "DROP TABLE " + s.tableName
	}

	if !isValidQuery(query) {
		query = ""
		SaveLog("", "Invalid query. Exiting...")
	}

	return query
}

func (s Schema) GetQueryString() string {
	return s.buildQuery()
}
