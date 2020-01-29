package interest

import (
	"github.com/sirupsen/logrus"
	"math"
	"time"
)

func DoSomeWork() {
	start := time.Now()
	var total float64 = 0
	for i := 1000000; i > 0; i-- {
		x := math.Sqrt(float64(i))
		y := x
		x = y
		total += x
	}
	logrus.Println(time.Since(start))
	logrus.Println(total)
}
