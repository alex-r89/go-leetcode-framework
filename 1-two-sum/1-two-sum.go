package leetcode

func twoSum(nums []int, target int) []int {
	indexMap := map[int]int{}

	for i, num := range nums {
		if idx, found := indexMap[target-num]; found {
			return []int{idx, i}
		}
		indexMap[num] = i
	}

	return nil
}
