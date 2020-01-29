package interest

import "net/http"

func NotifyDone() {
	http.Get("http://counter:8080/add")
}
