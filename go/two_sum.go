// https://leetcode.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	visited := make(map[int]int)

	for index, val := range nums {
		if _, ok := visited[target-val]; ok {
			return []int{int(visited[target-val]), int(index)}
		}
		visited[val] = index
	}
	return []int{-1, -1}
}