package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/google/uuid"
	"github.com/just1689/playing-with-hpa/model"
	"github.com/nsqio/go-nsq"
	"github.com/sirupsen/logrus"
	"net/http"
)

var addr = flag.String("address", ":8080", "The listen address for the batch server")
var nsqAddr = flag.String("nsqAddr", "127.0.0.1:4150", "Address and port of the nsq 4150 listener")

func main() {
	flag.Parse()

	logrus.Println("Starting batch-init!")
	logrus.Println("Address: ", *addr)
	logrus.Println("NSQd address:", *nsqAddr)

	http.HandleFunc("/", handleStartBatchRequest)
	http.ListenAndServe(*addr, nil)

}

func handleStartBatchRequest(w http.ResponseWriter, r *http.Request) {
	batchID := uuid.New().String()
	err := startBatch(*nsqAddr, batchID)
	if err != nil {
		http.Error(w, "failed to start batch", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(fmt.Sprint("ok: ", batchID)))

}

func startBatch(nsqAddr string, batchID string) (err error) {

	config := nsq.NewConfig()
	w, _ := nsq.NewProducer(nsqAddr, config)

	for i := 1; i <= 1000000; i++ {
		instruction := model.BatchInstruction{
			BatchID:   batchID,
			AccountID: fmt.Sprint("account-", i),
		}
		var b []byte
		b, err = json.Marshal(instruction)
		if err != nil {
			logrus.Errorln(err)
			return
		}
		err = w.Publish("batch", b)
		if err != nil {
			logrus.Errorln("Could not connect")
			return
		}
	}

	w.Stop()
	return

}
