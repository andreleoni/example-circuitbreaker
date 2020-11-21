package circuitbreaker

type BreakerIface interface {
	Wrap(work func() RateLimitServiceResponse)
	Opened() bool
	SemiOpened() bool
	Service() string
}

type StorageIface interface {
	AddToList(list_key, value string) error

	GetList(list_key string, size int64) ([]string, error)

	EraseList(list_key string) error

	Put(key string, value interface{}) error

	Get(key string) string
}
