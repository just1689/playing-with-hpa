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
	logrus.Println("Starting...")

	job := os.Getenv("job")
	addr := os.Getenv("address")
	nsqAddr := os.Getenv("nsqAddr")

	if job == "batch" {
		logrus.Println("> As batch...")
		logrus.Println("Address: ", addr)
		logrus.Println("NSQd address:", nsqAddr)
		batch.StartBatchServer(addr, nsqAddr)
	} else if job == "worker" {
		logrus.Println("> As worker...")
		logrus.Println("NSQd address:", nsqAddr)
		interest.StartInterestWorker(nsqAddr)

	}

	logrus.Errorln("Nothing left to do. Shutting down")

}
