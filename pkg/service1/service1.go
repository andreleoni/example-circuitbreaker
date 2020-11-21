package service1

import "example/pkg/circuitbreaker"

type Service1 struct {
	breaker circuitbreaker.BreakerIface
}

func New(breaker circuitbreaker.BreakerIface) *Service1 {
	return &Service1{breaker}
}

func (as Service1) CreateOrder() circuitbreaker.RateLimitServiceResponse {
	var serviceResponse circuitbreaker.RateLimitServiceResponse

	action := func() circuitbreaker.RateLimitServiceResponse {
		// Do something external
		serviceResponse = circuitbreaker.RateLimitServiceResponse{"ratelimit"}
		return serviceResponse
	}

	as.breaker.Wrap(action)

	return serviceResponse
}
