package main

func main() {

	println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
	println(findMaxAverage([]int{5}, 1))

}

func findMaxAverage(nums []int, k int) float64 {
	sum := 0

	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	maxSum := sum

	for i := k; i < len(nums); i++ {
		sum += nums[i]
		sum -= nums[i-k]
		if sum > maxSum {
			maxSum = sum
		}
	}

	return float64(maxSum) / float64(k)
}
