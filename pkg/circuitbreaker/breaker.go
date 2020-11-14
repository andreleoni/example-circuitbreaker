package circuitbreaker

import "fmt"

type Breaker struct {
	service       string
	errThreshould int64
}

func New(service string, errThreshould int64) *Breaker {
	return &Breaker{service, errThreshould}
}

func (b *Breaker) Wrap(work func() RateLimitServiceResponse) {
	if b.State() == "closed" {
		fmt.Println("CLOSED:", b.service)
	}

	fmt.Println("Calling service:", b.service)

	workResult := work()

	if workResult.RateLimit() {
		b.errThreshould++

		fmt.Println("Service error:", b.service, b.errThreshould)
	}
}

func (b *Breaker) State() string {
	if b.errThreshould > 1 {
		return "closed"
	}

	return "opened"
}
