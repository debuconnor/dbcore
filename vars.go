package dbcore

const (
	APP_NAME = "DBCORE"

	ERROR_DB_CONNECTION_FAILED        = "Failed to connect to database."
	ERROR_CREATE_SECRETMANAGER_CLIENT = "Failed to create SecretManager client."
	ERROR_ACCESS_SECRET_VERSION       = "Failed to access secret version."
	ERROR_DATA_CORRUPTION             = "Secret data corruption."
	ERROR_PARSE_JSON                  = "Failed to parse JSON."
	ERROR_INVALID_QUERY               = "Invalid query."
	ERROR_EXECUTING_QUERY             = "Failed to execute query."
)
