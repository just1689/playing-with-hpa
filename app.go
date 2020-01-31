package main

import (
	"flag"
	"github.com/just1689/playing-with-hpa/batch"
	"github.com/just1689/playing-with-hpa/counter"
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
		batch.UpdateBatchFromEnv()
		batch.StartBatchServer(addr, nsqAddr)
	} else if job == "worker" {
		logrus.Println("> As worker...")
		logrus.Println("NSQd address:", nsqAddr)
		interest.StartInterestWorker(nsqAddr)
	} else if job == "counter" {
		logrus.Println("> As counter...")
		logrus.Println("Address: ", addr)
		counter.StartCounter(addr)
	}

	logrus.Errorln("Nothing left to do. Shutting down")

}
