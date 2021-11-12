// python
// def twoSum(nums, target):
//
//	for i in nums:
//		for j in nums:
//			if nums[i] + nums[j] == target:
//				return i,j
package twoSum

func twoSum(nums []int, target int) []int {
	// m is hashmap
	m := make(map[int]int)
	for idx, num := range nums {
		if v, found := m[target-num]; found {
			return []int{v, idx}
		}
		m[num] = idx
	}
	return nil
}