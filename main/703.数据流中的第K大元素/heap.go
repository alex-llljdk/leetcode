package main

type minHeap struct {
	k    int   //capacity
	heap []int // heap数组
}

func createMinHeap(k int, nums []int) *minHeap {
	heap := &minHeap{k: k, heap: []int{}}
	for _, n := range nums {
		heap.add(n)
	}
	return heap
}

func (this *minHeap) add(num int) {
	if len(this.heap) < this.k {
		this.heap = append(this.heap, num)
		this.up(len(this.heap) - 1)
	} else if num > this.heap[0] {
		this.heap[0] = num
		this.down(0)
	}
}

func (this *minHeap) up(i int) { // 将索引i上的元素，上浮到合适位置
	for i > 0 {                  // 上浮到索引0就停止上浮
		parent := (i - 1) >> 1                // 找到父节点在heap数组中的位置
		if this.heap[parent] > this.heap[i] { // 如果父节点的数字比插入的数字大
			this.heap[parent], this.heap[i] = this.heap[i], this.heap[parent] // 交换
			i = parent                                                        // 更新 i
		} else { // 父比自己小，满足最小堆的性质，break
			break
		}
	}
}

作者：笨猪爆破组
链接：https://leetcode.cn/problems/kth-largest-element-in-a-stream/solutions/600672/zui-xiao-dui-de-mo-ban-dai-ma-zhi-jie-wa-hkvc/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

type KthLargest struct {
}

func Constructor(k int, nums []int) KthLargest {

}

func (this *KthLargest) Add(val int) int {

}

func main() {

}
