package corredor

import (
	"fmt"
	"net/http"
	"time"
)

var Secondslimit10 = 10 * time.Second

func Corredor(url1 string, url2 string) (string, error) {
	return Configurable(url1, url2, Secondslimit10)
}

func Configurable(url1 string, url2 string, limitTime time.Duration) (fastest string, erro error) {
	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(limitTime):
		return "", fmt.Errorf("time limit exceeded for %v and %v", url1, url2)
	}
}

func ping(url string) chan bool {
	channel := make(chan bool)
	go func() {
		http.Get(url)
		channel <- true
	}()

	return channel
}
