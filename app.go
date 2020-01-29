package main

import (
	"flag"
	"github.com/just1689/playing-with-hpa/batch"
	"github.com/sirupsen/logrus"
)

var job = flag.String("job", "batch", "Needs to be batch or worker")

var addr = flag.String("address", ":8080", "The listen address for the batch server")
var nsqAddr = flag.String("nsqAddr", "127.0.0.1:4150", "Address and port of the nsq 4150 listener")

func main() {
	flag.Parse()
	logrus.Println("Starting batch-init!")

	if *job == "batch" {
		logrus.Println("Address: ", *addr)
		logrus.Println("NSQd address:", *nsqAddr)
		batch.StartBatchServer(addr, nsqAddr)
	} else if *job == "worker" {
		logrus.Println("NSQd address:", *nsqAddr)

	}

}
