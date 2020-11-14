package circuitbreaker

type BreakerIface interface {
	Wrap(work func() RateLimitServiceResponse)
	State() string
}
