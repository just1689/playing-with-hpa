package interest

import (
	"context"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func NotifyDone() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	req, err := http.NewRequest("GET", "http://counter.default.svc.cluster.local:8080/add", nil)
	if err != nil {
		logrus.Errorln("Request error", err)
	}
	_, err = http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		logrus.Errorln("Do Request error", err)
	}
}
