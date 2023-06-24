# Task on O(n\*log(n))

#### En

Find two numbers in the array, the sum of which will be equal to the specified number

#### The solution

1. Iterate an array of numbers
2. Calculate a pair for each, subtracting the current number from the sum
3. Check that such a number is in the array using binary search

---

#### Ru

Найдите в массиве два числа, сумма которых будет равна указанному числу

### Решение

1. Итерируем массив чисел
2. Вычисляем для каждого пару, вычитая из суммы текущее число
3. Проверяем, что такое число есть в массиве с помощью бинарного поиска

---

### The solution on Golang

#### Run tests

`go test -v .`

```
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
```
