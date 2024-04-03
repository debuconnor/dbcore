package dbcore

const (
	APP_NAME = "DBCORE"

	ORDER_ASC          = "ASC"
	ORDER_DESC         = "DESC"
	EQUAL              = "="
	NOT_EQUAL          = "<>"
	GREATER_THAN       = ">"
	LESS_THAN          = "<"
	GREATER_THAN_EQUAL = ">="
	LESS_THAN_EQUAL    = "<="
	LIKE               = "LIKE"
	NOT_LIKE           = "NOT LIKE"
	IN                 = "IN"
	NOT_IN             = "NOT IN"
	BETWEEN            = "BETWEEN"
	NOT_BETWEEN        = "NOT BETWEEN"
	AND                = "AND"
	OR                 = "OR"
	INNER_JOIN         = "INNER JOIN"
	LEFT_JOIN          = "LEFT JOIN"
	RIGHT_JOIN         = "RIGHT JOIN"

	ERROR_DB_CONNECTION_FAILED        = "Failed to connect to database."
	ERROR_CREATE_SECRETMANAGER_CLIENT = "Failed to create SecretManager client."
	ERROR_ACCESS_SECRET_VERSION       = "Failed to access secret version."
	ERROR_DATA_CORRUPTION             = "Secret data corruption."
	ERROR_PARSE_JSON                  = "Failed to parse JSON."
	ERROR_INVALID_QUERY               = "Invalid query."
	ERROR_EXECUTING_QUERY             = "Failed to execute query."
)
