package main

//队列法
type MovingAverage struct {
	size  int
	sum   int
	queue []int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{
		size:  size,
		sum:   0,
		queue: []int{},
	}
}

func (this *MovingAverage) Next(val int) float64 {
	if len(this.queue) == this.size {
		this.sum -= this.queue[0]
		this.queue = this.queue[1:]
	}
	this.sum += val
	this.queue = append(this.queue, val)
	return float64(this.sum) / float64(len(this.queue))
}

//循环数组法
type MovingAverage struct {
	s   int
	cnt int
	arr []int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	arr := make([]int, size)
	return MovingAverage{
		0, 0, arr,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	idx := this.cnt % len(this.arr)
	this.s += val - this.arr[idx]
	this.arr[idx] = val
	this.cnt++
	return float64(this.s) / float64(min(this.cnt, len(this.arr)))
}
func main() {

}
