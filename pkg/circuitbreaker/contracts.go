package circuitbreaker

import "time"

type BreakerIface interface {
	Wrap(work func() RateLimitServiceResponse)
	Opened() bool
	SemiOpened() bool
	Service() string
}

type StorageIface interface {
	OpenCircuitList() []int

	ClearOpenCircuitList()

	AddSuccess()

	AddError()

	SetLastErrorOcurredAt()

	LastErrorOcurredAt() time.Time
}
