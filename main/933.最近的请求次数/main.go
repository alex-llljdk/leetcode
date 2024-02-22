package main

//队列法
type RecentCounter struct {
	cntqueue []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		cntqueue: []int{},
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.cntqueue = append(this.cntqueue, t)
	for i, v := range this.cntqueue {
		if v < t-3000 {
			continue
		}
		this.cntqueue = this.cntqueue[i:]
		break
	}
	return len(this.cntqueue)
}

func main() {

}
