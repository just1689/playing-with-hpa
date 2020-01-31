package counter

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/just1689/playing-with-hpa/model"
	"github.com/nsqio/go-nsq"
	"github.com/sirupsen/logrus"
	"net/http"
	"sync/atomic"
)

var ops uint64

func StartCounter(addr string, nsqAddr string) {

	go func() {
		config := nsq.NewConfig()
		config.MaxInFlight = 10
		q, _ := nsq.NewConsumer(model.BatchTopicName, fmt.Sprint(uuid.New().String(), "#ephemeral"), config)
		q.AddConcurrentHandlers(HandleMessage, 4)
		if err := q.ConnectToNSQD(nsqAddr); err != nil {
			logrus.Panic("Could not connect to NSQ for subscribe", nsqAddr)
		}
	}()

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		atomic.AddUint64(&ops, 1)
		w.Write([]byte("ok"))
	})
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		result := atomic.LoadUint64(&ops)
		w.Write([]byte(fmt.Sprint(result)))
	})
	http.ListenAndServe(addr, nil)
}

var HandleMessage nsq.HandlerFunc = func(message *nsq.Message) error {
	atomic.AddUint64(&ops, 1)
	return nil
}
