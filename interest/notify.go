package interest

import (
	"github.com/just1689/playing-with-hpa/model"
	"github.com/nsqio/go-nsq"
	"github.com/sirupsen/logrus"
)

func StartPublisher(nsqAddr string) {
	go func() {
		config := nsq.NewConfig()
		config.MaxInFlight = 10
		p, err := nsq.NewProducer(nsqAddr, config)
		if err != nil {
			logrus.Fatal(err)
		}
		NotifyDone = func() {
			//TODO: retry?
			err := p.Publish(model.CountTopicName, []byte(" "))
			if err != nil {
				logrus.Errorln(err)
			}
		}
	}()
}

var NotifyDone = func() {
	logrus.Fatalln("NOTIFY NOT IMPLEMENTED")
}
