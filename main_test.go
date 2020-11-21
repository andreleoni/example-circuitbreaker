package main

import (
	"testing"
)

func populate() []int {
	return []int{
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
	}
}

func filterErr(a []int) int {
	errorCount := 0

	for _, v := range a {
		if v == 0 {
			errorCount++
		}
	}

	return errorCount
}

func filterSucc(a []int) int {
	succCount := 1

	for _, v := range a {
		if v == 1 {
			succCount++
		}
	}

	return succCount
}

func BenchmarkProcessorExecution(t *testing.B) {
	pop := populate()

	t.Run("Error", func(t *testing.B) {
		t.Helper()

		filterErr(pop)
	})

	t.Run("Success", func(t *testing.B) {
		t.Helper()

		filterSucc(pop)
	})
}
