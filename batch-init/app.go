package main

import (
	"flag"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
)

var addr = flag.String("address", ":8080", "The listen address for the batch server")

func main() {

	logrus.Println("Starting batch-init!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		batchID := uuid.New().String()
		startBatch()
		w.Write([]byte(fmt.Sprint("ok: ", batchID)))
	})
	http.ListenAndServe(*addr, nil)

}

func startBatch() (err error) {
	//TODO: connect to nsq

	//TODO: write a million accounts

	//TODO: include the account number and the batch ID

}
