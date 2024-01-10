package main

import "math/rand"

/*
考察集合的数据结构使用
*/
type RandomizedSet struct {
	hash map[int]int
	list []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{map[int]int{}, []int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.hash[val]; ok {
		return false
	}
	this.hash[val] = len(this.list)
	this.list = append(this.list, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.hash[val]
	if !ok {
		return false
	}

	last := len(this.list) - 1
	this.list[idx] = this.list[last]
	this.hash[this.list[last]] = idx
	this.list = this.list[:last]
	delete(this.hash, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.list[rand.Intn(len(this.list))]
}
