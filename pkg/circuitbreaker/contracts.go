package circuitbreaker

type BreakerIface interface {
	Wrap(work func() RateLimitServiceResponse)
	State() string
	Opened() bool
	Service() string
}

type StorageIface interface {
}
