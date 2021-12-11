package hash

func Sum(nums []int, target int) []int {
	hmap := make(map[int]int)
	for i, a := range nums {
		if pos, ok := hmap[target-a]; ok {
			return []int{i, pos}
		} else {
			hmap[a] = i
		}
	}
	return nil
}
