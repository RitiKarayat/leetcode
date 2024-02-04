// https://leetcode.com/problems/contains-duplicate/

func containsDuplicate(nums []int) bool {

	uniq := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		if !uniq[nums[i]] {
			uniq[nums[i]] = true
			continue
		}
		return true
	}
	return false
}