package main

import (
	"flag"
	"fmt"
	"github.com/google/uuid"
	"github.com/just1689/playing-with-hpa/model"
	"github.com/sirupsen/logrus"
	"net/http"
	"github.com/nsqio/go-nsq"
)

var addr = flag.String("address", ":8080", "The listen address for the batch server")
var nsqAddr = flag.String("nsqAddr", "127.0.0.1:4150", "Address and port of the nsq 4150 listener")

func main() {
	flag.Parse()

	logrus.Println("Starting batch-init!")
	logrus.Println("Address: ", *addr)
	logrus.Println("NSQd address:", *nsqAddr)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		batchID := uuid.New().String()
		startBatch(*nsqAddr)
		w.Write([]byte(fmt.Sprint("ok: ", batchID)))
	})
	http.ListenAndServe(*addr, nil)

}




func startBatch(nsqAddr string) (err error) {

	i := model.BatchInstruction{
	}


	config := nsq.NewConfig()
	w, _ := nsq.NewProducer(nsqAddr, config)

	err = w.Publish("batch", []byte("test"))
	if err != nil {
		logrus.Errorln("Could not connect")
		return
	}

	w.Stop()

	//TODO: connect to nsq

	//TODO: write a million accounts

	//TODO: include the account number and the batch ID

}
