package dbcore

type MainQuery struct {
	TableName  string
	Columns    []string
	Conditions []Condition
}

type JoinQuery struct {
	TableName  string
	Columns    []string
	Conditions []Condition
	JoinType   string
}

type Condition struct {
	Column   string
	Operator string
	Value    interface{}
}
