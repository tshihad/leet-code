package array

func MaxProfit(prices []int) int {
	var sub int
	min := 99999
	for _, p := range prices {
		if p < min {
			min = p
		}
		if p-min > sub {
			sub = p - min
		}
	}

	return sub
}
