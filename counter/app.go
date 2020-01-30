package counter

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
)

var BatchCounters Batch = Batch{
	Mutex: sync.Mutex{},
	jobs:  make(map[string]*uint64),
}

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

func (b *Batch) GetSafely() (result [][]string) {
	b.Lock()
	defer b.Unlock()
	for x, y := range b.jobs {
		var line []string
		line = append(line, x)
		line = append(line, strconv.FormatUint(*y, 10))
		result = append(result, line)
	}
	return
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
		w.Write([]byte("ok"))
	})
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		for _, row := range BatchCounters.GetSafely() {
			for _, bit := range row {
				w.Write([]byte(fmt.Sprint(bit, " - ")))
			}
			w.Write([]byte(" <br />"))
		}
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
