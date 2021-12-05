package initialization

import (
	"os"
)

var file *os.File

func init() {
	var err error
	file, err = os.OpenFile("src/logs/db.logs", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return
	}
}
