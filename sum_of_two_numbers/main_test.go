package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNumbers(t *testing.T) {
	type args struct {
		nums []int
		sum  int
	}

	type testCase struct {
		name string
		args
		expected []int
	}

	testCases := []testCase{
		{
			name: "Тестирование GetNumbers",
			args: args{
				nums: []int{1, 2, 4, 6, 8, 10, 13, 15, 16, 18, 19, 21, 24, 27, 28},
				sum:  28,
			},
			expected: []int{1, 27},
		},
		{
			name: "Тестирование GetNumbers",
			args: args{
				nums: []int{1, 2, 4, 6, 8, 10, 13, 15, 16, 18, 19, 21, 24, 27, 28},
				sum:  100,
			},
			expected: []int{},
		},
		{
			name: "Тестирование GetNumbers",
			args: args{
				nums: []int{-1, -2, 0, 6, 8, 10, 13, 15, 16, 18, 19, 21, 24, 27, 28},
				sum:  100,
			},
			expected: []int{},
		},
	}

	var received []int
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			received = GetNumbers(testCase.args.nums, testCase.args.sum)

			assert.EqualValues(t, testCase.expected, received)
		})
	}
}

func TestBinarySearchNumber(t *testing.T) {
	type args struct {
		nums []int
		n    int
	}

	type testCase struct {
		name string
		args
		expected bool
	}

	testCases := []testCase{
		{
			name: "Тестирование BinarySearchNumber",
			args: args{
				nums: []int{1, 2, 4, 6, 8, 10, 13, 15, 16, 18, 19, 21, 24, 27, 28},
				n:    16,
			},
			expected: true,
		},
		{
			name: "Тестирование BinarySearchNumber",
			args: args{
				nums: []int{-1, -2, 0, 6, 8, 10, 13, 15, 16, 18, 19, 21, 24, 27, 28},
				n:    -2,
			},
			expected: true,
		},
		{
			name: "Тестирование BinarySearchNumber",
			args: args{
				nums: []int{1, 2, 4, 6, 8, 10, 13, 15, 16, 18, 19, 21, 24, 27, 28},
				n:    -1,
			},
			expected: false,
		},
	}

	var received bool
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			received = BinarySearchNumber(testCase.n, testCase.nums)
			assert.EqualValues(t, testCase.expected, received)
		})
	}
}
