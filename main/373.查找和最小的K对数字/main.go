package main

import "container/heap"

//对路归并 最小堆  时间O（M+K）logM   空间oM
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	n, m, ans := len(nums1), len(nums2), [][]int{}
	flag := n > m //判断两个数组的长度
	if flag {     //优先处理长度短的数组
		n, m, nums1, nums2 = m, n, nums2, nums1
	}
	if n > k { //如果n长于k，那么只需要记录k个优先序对
		n = k
	}
	pq := make(hp, n)
	//nlogm
	for i := 0; i < n; i++ { //记录以a序列为基的优先序对，避免后续在处理a+1，b时造成重复 加入 a+2，b和a+1，b+1
		pq[i] = []int{nums1[i] + nums2[0], i, 0}
	}
	//klogn
	heap.Init(&pq)
	for pq.Len() > 0 && len(ans) < k { //堆长度>0,且答案小于k
		poll := heap.Pop(&pq).([]int)
		a, b := poll[1], poll[2] //取得预序列下标
		if flag {                //计算答案
			ans = append(ans, []int{nums2[b], nums1[a]})
		} else {
			ans = append(ans, []int{nums1[a], nums2[b]})
		}
		if b < m-1 { //判断下标是否满足较长数组
			heap.Push(&pq, []int{nums1[a] + nums2[b+1], a, b + 1})
		}
	}
	return ans
}

// 最小堆模板
type hp [][]int

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i][0] < h[j][0] }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.([]int)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func main() {

}
