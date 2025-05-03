package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElements(t *testing.T) {
	testCases := []struct {
		size    int
		isNil   bool
		isEmpty bool
	}{
		{size: -5, isNil: true},
		{size: 0, isNil: true},
		{size: 5, isNil: false, isEmpty: false},
	}

	for i, testCase := range testCases {
		res := generateRandomElements(testCase.size)
		if testCase.isNil {
			assert.Nil(t, res, "test case: %d", i)
		} else {
			assert.NotNil(t, res, "test case: %d", i)
			if !testCase.isEmpty {
				assert.NotEmpty(t, res, "test case: %d", i)
				assert.Len(t, res, testCase.size, "test case: %d", i)
			}
		}
	}
}

func TestMaximum(t *testing.T) {
	testCases := []struct {
		data       []int
		isEmpty    bool
		maxElement int
	}{
		{data: nil, isEmpty: true},
		{data: []int{}, isEmpty: true},
		{data: []int{1}, isEmpty: false, maxElement: 1},
		{data: []int{1, 2, 3, 4, 5}, isEmpty: false, maxElement: 5},
		{data: []int{-1, -2, -3, -4, -5}, isEmpty: false, maxElement: -1},
		{data: []int{-1, -2, -3, -4, -5, 0, 1, 2, 3, 4, 5}, isEmpty: false, maxElement: 5},
	}
	for i, testCase := range testCases {
		res := maximum(testCase.data)

		if testCase.isEmpty {
			assert.Empty(t, res, "test case: %d", i)
		} else {
			assert.NotEmpty(t, res, "test case: %d", i)
			assert.Equal(t, testCase.maxElement, res, "test case: %d", i)
		}
	}
}
