package circuitbreaker

import "fmt"

type Breaker struct {
	storage       StorageIface
	service       string
	errThreshould int64
	errCount      int64
}

func New(storage StorageIface, service string, errThreshould int64) *Breaker {
	return &Breaker{storage: storage, service: service, errThreshould: errThreshould}
}

func (b *Breaker) Wrap(work func() RateLimitServiceResponse) {
	if !b.Opened() {
		fmt.Println("CLOSED:", b.service)
		return
	}

	fmt.Println("Calling service:", b.service)

	workResult := work()

	if workResult.RateLimit() {
		b.errCount++

		fmt.Println("Service error:", b.service, b.errThreshould)
	}
}

func (b *Breaker) State() string {
	if b.errThreshould >= b.errCount {
		return "closed"
	}

	return "opened"
}

func (b *Breaker) Opened() bool {
	return b.State() == "opened"
}

func (b *Breaker) Service() string {
	return b.service
}
