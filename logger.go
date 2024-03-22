package dbcore

import "log"

func Log(msg string) {
	msg = "DBCORE: " + msg
	log.Println(msg)
}
