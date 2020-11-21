package circuitbreaker

import (
	"fmt"
	"time"
)

const listSize = 100

type states int

const (
	opened     states = iota
	closed     states = iota
	semiopened states = iota
)

const baseKey = "breakerstorage/redis/"
const baseKeyLastError = "breakerstorage/redis/last_error_ocurred_at/"

const redisKeyDateFormat = time.RFC3339

type bucket struct {
	breaker            Breaker
	list               *[]string
	lastErrorOccuredAt *time.Time
}

func (b *bucket) ConsecutiveSuccess() int {
	return 1
}

func (b *bucket) ErrPercentage() int {
	list := b.getList()

	successCount := 0
	errCount := 0

	for _, item := range *list {
		if item == "0" {
			errCount++
		}

		successCount++
	}

	return errCount / successCount
}

func (b *bucket) OpenCircuitList() []string {
	return *b.list
}

func (b *bucket) ClearOpenCircuitList() {
	key := fmt.Sprint(baseKey, b.breaker.service)

	b.breaker.storage.EraseList(key)
}

func (b *bucket) AddSuccess() {
	key := fmt.Sprint(baseKey, b.breaker.service)

	b.breaker.storage.AddToList(key, "1")
}

func (b *bucket) AddError() {
	key := fmt.Sprint(baseKey, b.breaker)

	b.breaker.storage.AddToList(key, "0")
}

func (b *bucket) SetLastErrorOcurredAt() {
	key := fmt.Sprint(baseKeyLastError, b.breaker)

	b.breaker.storage.Put(key, time.Now().Format(redisKeyDateFormat))
}

func (b *bucket) LastErrorOcurredAt() time.Time {
	key := fmt.Sprint(baseKeyLastError, b.breaker.service)

	timeFromStorage := b.breaker.storage.Get(key)

	var returnTime time.Time

	returnTime, err := time.Parse(redisKeyDateFormat, timeFromStorage)
	if err != nil {
		returnTime = time.Now()
		fmt.Println(err)
	}

	return returnTime
}

func (b *bucket) Opened() bool {
	return b.state() == opened
}

func (b *bucket) SemiOpened() bool {
	return b.state() == semiopened
}

func (b *bucket) state() states {
	if b.breaker.volumeThreshould > len(b.OpenCircuitList()) {
		if b.ErrPercentage() < b.breaker.errThreshould {
			if b.LastErrorOcurredAt().Add(b.breaker.sleepWindow).Before(time.Now()) {
				return semiopened
			}

			return closed
		}
	}

	return opened
}

func (b *bucket) getList() *[]string {
	if b.list != nil {
		return b.list
	}

	key := fmt.Sprint(baseKey, b.breaker)

	list, err := b.breaker.storage.GetList(key, 100)
	fmt.Println("erro ao buscar lista: ", err)

	b.list = &list

	return b.list
}
