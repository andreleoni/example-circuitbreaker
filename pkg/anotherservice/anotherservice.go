package anotherservice

import "example/pkg/circuitbreaker"

type AnotherService struct {
	breaker circuitbreaker.BreakerIface
}

func New(breaker circuitbreaker.BreakerIface) *AnotherService {
	return &AnotherService{breaker}
}

func (as AnotherService) CreateOrder() circuitbreaker.RateLimitServiceResponse {
	var serviceResponse circuitbreaker.RateLimitServiceResponse

	action := func() circuitbreaker.RateLimitServiceResponse {
		// Do something external
		serviceResponse = circuitbreaker.RateLimitServiceResponse{"ratelimit"}
		return serviceResponse
	}

	as.breaker.Wrap(action)

	return serviceResponse
}
