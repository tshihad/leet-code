package array

func productExceptSelf(nums []int) []int {
	prodArr := make([]int, len(nums))
	prod := 1
	zeroIndex := -1
	for i, n := range nums {
		if n == 0 {
			if zeroIndex >= 0 {
				return prodArr
			}
			zeroIndex = i
			continue
		} else {
			prod *= n
		}
	}
	if zeroIndex >= 0 {
		prodArr[zeroIndex] = prod
		return prodArr
	}
	for i, n := range nums {
		prodArr[i] = prod / n
	}

	return prodArr
}
