package crema

import (
	"fmt"
	"log"
	"os"
	"time"
)

var file *os.File
var err error

const fileName = "./api.log"

func InitLogFiles() {
	err = openFile()

	if err != nil {
		if err.Error() == "open "+fileName+": no such file or directory" {
			file, err = os.Create(fileName)
			handleError(err)
		}
	}
}

func openFile() error {
	file, err = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0600)

	return err
}

func closeFile() error {
	return file.Close()
}

func Printf(data interface{}) {
	file.WriteString(fmt.Sprintf("%v %v\n", time.Now().Format("2006-01-02 15:04:05"), data))
}

func PrintfError(data interface{}) {
	file.WriteString(fmt.Sprintf("%v %v\n", time.Now().Format("2006-01-02 15:04:05 [ERROR]"), data))
}

func LogPrintf(data interface{}) {
	log.Println(data)
	Printf(data)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
