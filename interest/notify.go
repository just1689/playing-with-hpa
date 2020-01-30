package interest

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func NotifyDone(batchID string) {
	ctx, cancel := context.WithTimeout(context.Background(), 75*time.Millisecond)
	defer cancel()
	req, err := http.NewRequest("GET", fmt.Sprint("http://counter:8080/add?batchID=", batchID), nil)
	if err != nil {
		logrus.Errorln("Request error", err)
	}
	_, err = http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		logrus.Errorln("Do Request error", err)
	}
}
