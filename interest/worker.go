package interest

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/just1689/playing-with-hpa/model"
	"github.com/nsqio/go-nsq"
	"github.com/sirupsen/logrus"
)

func StartInterestWorker(nsqAddr string) {
	go func() {
		config := nsq.NewConfig()
		config.MaxInFlight = 10
		q, _ := nsq.NewConsumer(model.BatchTopicName, fmt.Sprint(uuid.New().String(), "#ephemeral"), config)
		q.AddConcurrentHandlers(HandleMessage, 4)
		if err := q.ConnectToNSQD(nsqAddr); err != nil {
			logrus.Panic("Could not connect to NSQ for subscribe", nsqAddr)
		}
	}()
	select {}
}

type MessageHandler struct {
	MaxInFlight int
}

var HandleMessage nsq.HandlerFunc = func(message *nsq.Message) error {
	instruction := model.BatchInstruction{}
	err := json.Unmarshal(message.Body, &instruction)
	if err != nil {
		logrus.Errorln(err)
	}
	logrus.Println("Handling instruction: batch #", instruction.BatchID, " on account: ", instruction.AccountID)
	DoSomeWork()
	NotifyDone()
	return nil
}
