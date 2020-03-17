// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"project-hydra/Codes/Section_4/Hydra/hlogger"
// )

// func main() {
// 	logger := hlogger.GetInstance()
// 	logger.Println("Starting Hydra web service")

// 	http.HandleFunc("/", sroot)
// 	http.ListenAndServe(":8080", nil)
// }

// func sroot(w http.ResponseWriter, r *http.Request) {
// 	logger := hlogger.GetInstance()
// 	fmt.Fprint(w, "Welcome to the Hydra software system")

// 	logger.Println("Received an http Get request on root url")
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
