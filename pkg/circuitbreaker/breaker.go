package circuitbreaker

import (
	"fmt"
	"time"
)

type Breaker struct {
	storage                      StorageIface
	service                      string
	sleepWindow                  time.Duration
	errThreshould                int
	volumeThreshould             int
	consecutiveSuccessThreshould int
}

func New(
	storage StorageIface, service string, errThreshould int, sleepWindow time.Duration, consecutiveSuccessThreshould int) *Breaker {

	return &Breaker{storage: storage,
		service:                      service,
		errThreshould:                errThreshould,
		consecutiveSuccessThreshould: consecutiveSuccessThreshould,
		sleepWindow:                  sleepWindow}
}

func (b *Breaker) Wrap(work func() RateLimitServiceResponse) {
	currentBucket := bucket{breaker: *b}

	if !currentBucket.Opened() || !currentBucket.SemiOpened() {
		fmt.Println("CLOSED:", b.service)
		return
	}

	fmt.Println("Calling service:", b.service)

	workResult := work()

	if workResult.RateLimit() {
		currentBucket.AddError()
		currentBucket.SetLastErrorOcurredAt()
	} else if currentBucket.SemiOpened() {
		currentBucket.AddSuccess()

		if currentBucket.ConsecutiveSuccess() > b.consecutiveSuccessThreshould {
			currentBucket.ClearOpenCircuitList()
		}
	} else {
		currentBucket.AddSuccess()
	}
}

func (b *Breaker) Service() string {
	return b.service
}

func (b *Breaker) Opened() bool {
	currentBucket := bucket{breaker: *b}

	return currentBucket.Opened()
}

func (b *Breaker) SemiOpened() bool {
	currentBucket := bucket{breaker: *b}

	return currentBucket.SemiOpened()
}
