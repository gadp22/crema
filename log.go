// Copyright 2019 The Crema Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package crema

import (
	"fmt"
	"log"
	"os"
	"time"
)

var file *os.File
var err error

//TO DO: use config file
const fileName = "./api.log"

// InitLogFiles initializes the log file
// The log file will be placed under GOPATH/ named api.log
func InitLogFiles() {
	err = openFile()

	if err != nil {
		if err.Error() == "open "+fileName+": no such file or directory" {
			file, err = os.Create(fileName)

			if err != nil {
				panic(err)
			}
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

// Printf calls file.WriteString() to print log messages to the log file
// Example output: YYYY-MM-DD HH:MM:SS this is a message log.
func Printf(data interface{}) {
	file.WriteString(fmt.Sprintf("%v %v\n", time.Now().Format("2006-01-02 15:04:05"), data))
}

// PrintfError calls file.WriteString() to print error messages to the log file
// Example output: YYYY-MM-DD HH:MM:SS [ERROR] this is an error log.
func PrintfError(data interface{}) {
	file.WriteString(fmt.Sprintf("%v %v\n", time.Now().Format("2006-01-02 15:04:05 [ERROR]"), data))
}

// LogPrintf calls Printf to print to the log file
// It also calls log.Println() to print log messages to the standard logger
// Example output: YYYY-MM-DD HH:MM:SS this is a message log.
func LogPrintf(data interface{}) {
	log.Println(data)
	Printf(data)
}

// LogPrintfError calls file.WriteString() to print error messages to the log file
// It also calls log.Println() to print log messages to the standard logger
// Example output: YYYY-MM-DD HH:MM:SS [ERROR] this is an error log.
func LogPrintfError(data interface{}) {
	log.Println(data)
	PrintfError(data)
}
