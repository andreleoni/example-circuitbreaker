package circuitbreaker

type RateLimitServiceResponse struct {
	ErrorMsg string
}

func (rlsr *RateLimitServiceResponse) String() string {
	return rlsr.ErrorMsg
}

func (rlsr *RateLimitServiceResponse) RateLimit() bool {
	return rlsr.ErrorMsg == "ratelimit"
}
