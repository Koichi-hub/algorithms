package main

func GetNumbers(nums []int, sum int) []int {
	var n int
	for _, v := range nums {
		n = sum - v
		if exist := BinarySearchNumber(n, nums); exist {
			return []int{v, n}
		}
	}
	return []int{}
}

func BinarySearchNumber(n int, nums []int) bool {
	var mid, guess int
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid = (right + left) / 2
		guess = nums[mid]
		if n == guess {
			return true
		}
		if guess > n {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
