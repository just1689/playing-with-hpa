package counter

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

var ops uint64

func StartCounter(addr string) {
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		atomic.AddUint64(&ops, 1)
	})
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		result := atomic.LoadUint64(&ops)
		w.Write([]byte(fmt.Sprint(result)))
	})
	http.ListenAndServe(addr, nil)
}
