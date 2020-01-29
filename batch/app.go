package batch

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/just1689/playing-with-hpa/model"
	"github.com/nsqio/go-nsq"
	"github.com/sirupsen/logrus"
	"net/http"
)

func StartBatchServer(addr *string, nsqAddr *string) {
	http.HandleFunc("/", createStartBatchHandler(nsqAddr))
	http.ListenAndServe(*addr, nil)
}

func createStartBatchHandler(nsqAddr *string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		batchID := uuid.New().String()
		err := startBatch(*nsqAddr, batchID)
		if err != nil {
			http.Error(w, "failed to start batch", http.StatusInternalServerError)
			return
		}
		w.Write([]byte(fmt.Sprint("ok: ", batchID)))

	}
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
		err = w.Publish(model.BatchTopicName, b)
		if err != nil {
			logrus.Errorln("Could not connect")
			return
		}
	}

	w.Stop()
	return

}
