package main

import "github.com/emirpasic/gods/trees/redblacktree"

//树集合，红黑树,需要二分查找时间复杂度为O（nlogn） n为日程安排数量 空间复杂度为on
type MyCalendar struct {
	rbt *redblacktree.Tree
}

func Constructor() MyCalendar {
	return MyCalendar{
		rbt: redblacktree.NewWithIntComparator(),
	}
}

func (this *MyCalendar) Book(start int, end int) bool {
	if p, ok := this.rbt.Floor(start); ok && p.Value.(int) > start {
		return false
	}
	if p, ok := this.rbt.Ceiling(start); ok && p.Key.(int) < end {
		return false
	}
	this.rbt.Put(start, end)
	return true
}

func main() {

}
