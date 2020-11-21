package circuitbreaker

import (
	"fmt"
	"time"
)

type states int

const (
	opened     states = iota
	closed     states = iota
	semiopened states = iota
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
	currentBucket := bucket{b.storage}

	if !bucket.Opened() || !bucket.SemiOpened() {
		fmt.Println("CLOSED:", b.service)
		return
	}

	fmt.Println("Calling service:", b.service)

	workResult := work()

	if workResult.RateLimit() {
		bucket.AddError()
		bucket.SetLastErrorOcurredAt()
	} else if b.SemiOpened() {
		bucket.AddSuccess()

		if bucket.ConsecutiveSuccess() > b.consecutiveSuccessThreshould {
			bucket.ClearOpenCircuitList()
		}
	} else {
		bucket.AddSuccess()
	}
}

func (b *Breaker) Service() string {
	return b.service
}

func (b *Breaker) Opened() bool {
	return b.state() == opened
}

func (b *Breaker) SemiOpened() bool {
	return b.state() == semiopened
}

func (b *Breaker) state() states {
	if b.volumeThreshould > len(b.storage.OpenCircuitList()) {
		if b.storage.ErrPercentage() < b.errThreshould {
			if b.storage.LastErrorOcurredAt().Add(b.sleepWindow).Before(time.Now()) {
				return semiopened
			}

			return closed
		}
	}

	return opened
}
