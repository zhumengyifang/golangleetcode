package main

import (
	"fmt"
	"goleetcode/src/main/Model"
)

func testSortList() {
	p := Model.ListNode{Val: 4}
	p.Next = &Model.ListNode{Val: 2}
	p.Next.Next = &Model.ListNode{Val: 1}
	p.Next.Next.Next = &Model.ListNode{Val: 3}

	result := SortList(&p)
	fmt.Print(result)
}

func SortList(head *Model.ListNode) *Model.ListNode {
	if head == nil {
		return nil
	}
	return mergeSort(head)
}

func mergeSort(head *Model.ListNode) *Model.ListNode {
	if head.Next == nil {
		return head
	}
	p := head
	q := head
	var per *Model.ListNode
	for ; q != nil && q.Next != nil; {
		per, p, q = p, p.Next, q.Next.Next
	}
	per.Next = nil
	l := mergeSort(head)
	r := mergeSort(p)
	return merge(l, r)
}

func merge(l *Model.ListNode, r *Model.ListNode) *Model.ListNode {
	dummyHead := &Model.ListNode{}
	cur := dummyHead
	for ; l != nil && r != nil; {
		if l.Val <= r.Val {
			cur.Next = l
			cur = cur.Next
			l = l.Next
		} else {
			cur.Next = r
			cur = cur.Next
			r = r.Next
		}
	}
	if l != nil {
		cur.Next = l
	}
	if r != nil {
		cur.Next = r
	}
	return dummyHead.Next
}
