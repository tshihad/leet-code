package array

func maxSubArrayBruteForce(nums []int) int {
	m := -999999
	for i := range nums {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			m = max(m, sum)
		}
	}

	return m
}

func maxSubArrayRecursive(nums []int) int {
	return solve(nums, 0, false)
}

func solve(A []int, i int, mustPick bool) int {
	if i >= len(A) {
		if mustPick {
			return 0
		}
		return -99999
	}

	if mustPick {
		return max(0, A[i]+solve(A, i+1, true))
	}
	return max(solve(A, i+1, false), A[i]+solve(A, i+1, true))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
