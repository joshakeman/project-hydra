// package hlogger

// import (
// 	"log"
// 	"os"
// 	"sync"
// )

// type hydraLogger struct {
// 	*log.Logger
// 	filename string
// }

// var hlogger *hydraLogger
// var once sync.Once

// //GetInstance create a singleton instance of the hydra logger
// func GetInstance() *hydraLogger {
// 	once.Do(func() {
// 		hlogger = createLogger("hydralogger.log")
// 	})
// 	return hlogger
// }

// //Create a logger instance
// func createLogger(fname string) *hydraLogger {
// 	file, _ := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)

// 	return &hydraLogger{
// 		filename: fname,
// 		Logger:   log.New(file, "Hydra ", log.Lshortfile),
// 	}
// }

package main

import (
	"encoding/csv"
	"log"
	"os"
)

var data = [][]string{{"Line1", "Hello Readers of"}, {"Line2", "golangcode.com"}}

func main() {
	file, err := os.Create("result.csv")
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
