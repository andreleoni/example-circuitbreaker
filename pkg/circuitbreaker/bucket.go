package storagememoizer

import "time"

type bucket struct {
	storage            StorageIface
	list               *[]string
	lastErrorOccuredAt *time.Time
}

func (b *bucket) ConsecutiveSuccess() int {
	return 1
}

func (b *bucket) ErrPercentage() int {
	return 1
}

func (b *bucket) getList() {}
