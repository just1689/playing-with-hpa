package counter

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"sync"
	"sync/atomic"
)

var BatchCounters Batch

type Batch struct {
	sync.Mutex
	jobs map[string]*uint64
}

func (b *Batch) AddSafely(name string) {
	b.Lock()
	defer b.Unlock()
	counter, found := b.jobs[name]
	if found {
		atomic.AddUint64(counter, 1)
	} else {
		var counter *uint64
		b.jobs[name] = counter
		atomic.AddUint64(counter, 1)
	}
}

func (b *Batch) GetSafely(name string) uint64 {
	b.Lock()
	defer b.Unlock()
	counter, found := b.jobs[name]
	if found {
		return *counter
	} else {
		logrus.Println("could not get safely - no batch by name")
	}
	return 0
}

func StartCounter(addr string) {
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		batchID, err := getBatchID(r)
		if err != nil {
			logrus.Errorln("batchID is missing")
			http.Error(w, "no batchID supplied", http.StatusBadRequest)
			return
		}
		BatchCounters.AddSafely(batchID)
		i := BatchCounters.GetSafely(batchID)
		w.Write([]byte(fmt.Sprint(i)))
	})
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		batchID, err := getBatchID(r)
		if err != nil {
			logrus.Errorln("batchID is missing")
			http.Error(w, "no batchID supplied", http.StatusBadRequest)
			return
		}
		i := BatchCounters.GetSafely(batchID)
		w.Write([]byte(fmt.Sprint(i)))
	})
	http.ListenAndServe(addr, nil)
}

func getBatchID(r *http.Request) (string, error) {
	keys, ok := r.URL.Query()["batchID"]
	if !ok || len(keys[0]) < 1 {
		return "", errors.New("no batchID supplied")
	}
	batchID := keys[0]
	return batchID, nil
}
