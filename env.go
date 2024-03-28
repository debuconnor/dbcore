package dbcore

import "os"

func getEnv(envName string) string {
	return os.Getenv(envName)
}
