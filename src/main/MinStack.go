package main

import "math"

type MinStack struct {
	minNum int
	nums   []int
	index  int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{minNum: math.MaxInt32, nums: []int{}, index: 0}
}

func (this *MinStack) Push(x int) {
	if x < this.minNum {
		this.minNum = x
	}
	this.index++
	this.nums = append(this.nums, x)
}

func (this *MinStack) Pop() {
	if this.minNum == this.Top() {
		this.minNum = math.MaxInt32
	}

	this.index--
	this.nums = this.nums[0:this.index]

	for x := range this.nums {
		if this.nums[x] < this.minNum {
			this.minNum = this.nums[x]
		}
	}
}

func (this *MinStack) Top() int {
	return this.nums[this.index-1]
}

func (this *MinStack) GetMin() int {
	return this.minNum
}
