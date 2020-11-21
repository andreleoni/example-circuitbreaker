package service2

import "example/pkg/circuitbreaker"

type Service2 struct {
	breaker circuitbreaker.BreakerIface
}

func New(breaker circuitbreaker.BreakerIface) *Service2 {
	return &Service2{breaker}
}

func (os *Service2) DoSomething() circuitbreaker.RateLimitServiceResponse {
	var serviceResponse circuitbreaker.RateLimitServiceResponse

	action := func() circuitbreaker.RateLimitServiceResponse {
		// Do something external
		serviceResponse = circuitbreaker.RateLimitServiceResponse{"ratelimit"}
		return serviceResponse
	}

	os.breaker.Wrap(action)

	return serviceResponse
}
