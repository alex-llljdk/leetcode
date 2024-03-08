package main

func getID(x, w int) int {
	if x >= 0 {
		return x / w
	}

	return (x+1)/w - 1
}

// 桶排序 边界即为桶的大小，通过，哈希判断   时间复杂度on  空间复杂度o（min（n，indexDiff）） 最多有indexdiff个空间，或最长数组长度
func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	mp := map[int]int{}
	for i, x := range nums {
		id := getID(x, valueDiff+1)
		if _, has := mp[id]; has {
			return true
		}
		if y, has := mp[id+1]; has && abs(x-y) <= valueDiff {
			return true
		}
		if y, has := mp[id-1]; has && abs(x-y) <= valueDiff {
			return true
		}
		mp[id] = x
		if i >= indexDiff {
			delete(mp, getID(nums[i-indexDiff], valueDiff+1))
		}
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {

}
