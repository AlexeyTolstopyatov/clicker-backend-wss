package Handling

import (
	"errors"
	"log"
	"math"
	"time"
)

var (
	// RequestDeltaTimeError
	// Малый интервал между запросами
	RequestDeltaTimeError = errors.New("too many same requests")

	// requestStack
	// Хранилище сегментов
	requestStack = RequestStack{}

	// requestIndex
	// Указатель на сегмент
	requestIndex = 0
)

type RequestStack struct {
	first  int64
	second int64
}

func FixRequestTime() {
	if requestIndex%2 == 0 {
		requestStack.first = time.Now().UnixMilli()
	} else {
		requestStack.second = time.Now().UnixMilli()
	}
	requestIndex += 1
	logRequestStack()
}

func IsFirstRequest() bool {
	return requestIndex == 0
}

func logRequestStack() {
	log.Printf("Delta: %d\n", requestStack.second-requestStack.first)
}

func ProcessRequestTime() error {
	if math.Abs(float64(requestStack.second-requestStack.first)) < 1000 {
		requestIndex = 0
		requestStack = RequestStack{}
		return RequestDeltaTimeError
	}

	return nil
}
