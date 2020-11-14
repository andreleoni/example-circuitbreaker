package oneservice

import "example/pkg/circuitbreaker"

type OneService struct {
	breaker circuitbreaker.BreakerIface
}

func New(breaker circuitbreaker.BreakerIface) *OneService {
	return &OneService{breaker}
}

func (os *OneService) DoSomething() circuitbreaker.RateLimitServiceResponse {
	var serviceResponse circuitbreaker.RateLimitServiceResponse

	action := func() circuitbreaker.RateLimitServiceResponse {
		// Do something external
		serviceResponse = circuitbreaker.RateLimitServiceResponse{"ratelimit"}
		return serviceResponse
	}

	os.breaker.Wrap(action)

	return serviceResponse
}
