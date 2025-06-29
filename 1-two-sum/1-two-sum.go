package leetcode

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
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
