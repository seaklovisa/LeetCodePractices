package main

import (
	"fmt"
	"math"
)

func main() {
	//nums := []int{10, -10}
	//fmt.Print(minMax(nums))
	//fmt.Print(Average(nums))
	//fmt.Print(Reverse("hello"))
	//fmt.Print(Sum(0))
	//fmt.Print(IsPrime(3))
	//fmt.Print(Fib(5))
	fmt.Print(FactorialIter(5))
	fmt.Print(FactorialRec(6))
}

func minMax(nums []int) (min, max int) {
	if len(nums) == 0 {
		panic("empty array")
	}

	min, max = nums[0], nums[0]

	for _, n := range nums[1:] {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max
}

func Average(nums []int) float64 {

	if len(nums) == 0 {
		return 0
	}

	total := 0
	for _, n := range nums {
		total += n
	}

	return float64(total) / float64(len(nums))
}

func Reverse(s string) string {

	rs := []rune(s)

	for l, r := 0, len(rs)-1; l < r; l, r = l+1, r-1 {
		rs[l], rs[r] = rs[r], rs[l]
	}

	return string(rs)
}

/*
// 避免溢位：先把 n 或 n+1 中的偶數除以 2

	if n % 2 == 0 {
	    return (n / 2) * (n + 1)
	} else {
		//奇數 + (n+1) = 偶數
	    return n * ((n + 1) / 2)
	}
*/
func Sum(n int) int {
	if n < 0 {
		return 0
	}
	return n * (n + 1) / 2
}

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	//2是質數內唯一的偶數
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	i := 3
	end := int(math.Sqrt(float64(n)))

	for i <= end {
		if n%i == 0 {
			return false
		}
		//每次用奇數去判斷
		i += 2
	}

	return true
}

func Fib(n int) []int {
	result := make([]int, 0, n)
	a, b := 0, 1

	for i := 0; i < n; i++ {
		result = append(result, a)
		a, b = b, a+b
	}

	return result
}

// 階層迴圈版
func FactorialIter(n int) int {
	if n < 0 {
		return 0
	}

	res := 1

	for i := 2; i <= n; i++ {
		res *= i
	}

	return res
}

// 階層遞迴版
func FactorialRec(n int) int {
	if n < 0 {
		return 0
	}

	FactorialRec(n - 1)

}
