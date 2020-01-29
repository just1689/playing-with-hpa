package main

import (
	"flag"
	"github.com/just1689/playing-with-hpa/batch"
	"github.com/just1689/playing-with-hpa/interest"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	flag.Parse()
	logrus.Println("Starting batch-init!")

	job := os.Getenv("job")
	addr := os.Getenv("address")
	nsqAddr := os.Getenv("nsqAddr")

	if job == "batch" {
		logrus.Println("Address: ", addr)
		logrus.Println("NSQd address:", nsqAddr)
		batch.StartBatchServer(addr, nsqAddr)
	} else if job == "worker" {
		logrus.Println("NSQd address:", nsqAddr)
		interest.StartInterestWorker(nsqAddr)
		select {}

	}

	logrus.Errorln("Nothing left to do. Shutting down")

}
