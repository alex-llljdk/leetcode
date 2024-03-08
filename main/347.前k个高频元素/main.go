package main

func u2d(nums []int, index int, m map[int]int) {
	dad := index
	son := dad*2 + 1
	for son < len(nums) {
		if son+1 < len(nums) && m[nums[son+1]] < m[nums[son]] {
			son++
		}
		if m[nums[dad]] < m[nums[son]] {
			return
		}
		nums[dad], nums[son] = nums[son], nums[dad]
		dad = son
		son = dad*2 + 1
	}
}
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	var arr []int
	for i := range nums {
		if _, ok := m[nums[i]]; ok {
			m[nums[i]]++
		} else {
			m[nums[i]] = 1
			arr = append(arr, nums[i])
		}
	}

	for i := (k+1)/2 - 1; i >= 0; i-- {
		u2d(arr[:k], i, m)
	}

	for i := k; i < len(arr); i++ {
		if m[arr[i]] > m[arr[0]] {
			arr[0] = arr[i]
			u2d(arr[:k], 0, m)
		}
	}
	return arr[:k]

}

func main() {

}
