package dbcore

import "log"

func Log(msg ...interface{}) {
	if !IS_DEBUG {
		return
	}
	msg = append([]interface{}{"DBCORE: "}, msg...)
	log.Println(msg...)
}
