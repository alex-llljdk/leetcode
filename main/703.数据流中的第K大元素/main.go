package main

type KthLargest struct {
	nums []int
	k    int
}

type IntHeap []int

var h *IntHeap

func QuickSort(array []int, left int, right int) {
	if len(array) < 2 {
		return
	}
	cur, lo := left+1, left
	for cur < right {
		if array[cur] < array[left] {
			array[lo+1], array[cur] = array[cur], array[lo+1]
			lo++
		}
		cur++
	}
	array[left], array[lo] = array[lo], array[left]
	QuickSort(array, left, lo-1)
	QuickSort(array, lo+1, right)
}

func BinarySearch(nums []int, val int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] == val {
			return mid
		}
		if nums[mid] < val {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func Constructor(k int, nums []int) KthLargest {
	QuickSort(nums, 0, len(nums)-1)
	return KthLargest{nums, k}
}

func (this *KthLargest) Add(val int) int {
	idx := BinarySearch(this.nums, val)
	tmp1 := this.nums[:idx+1]
	tmp2 := this.nums[idx:]
	this.nums = append(tmp1, val)
	this.nums = append(this.nums, tmp2)
}

func main() {

}
