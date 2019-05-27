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

	result := SortListNode(&p)
	fmt.Print(result)
}

/**
排序链表
*/
func SortListNode(head *Model.ListNode) *Model.ListNode {
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

/**
回文链表
*/
func isPalindrome(head *Model.ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	fast := head //快指针
	slow := head //慢指针

	for ; fast.Next != nil && fast.Next.Next != nil; {
		fast = fast.Next.Next
		slow = slow.Next
	}

	slow = reverse(slow.Next)

	for ; slow != nil; {
		if head.Val != slow.Val {
			return false
		}
		head = head.Next
		slow = slow.Next
	}
	return true
}

/**
反转链表
*/
func reverse(node *Model.ListNode) *Model.ListNode {
	if node.Next == nil {
		return node
	}
	newHead := reverse(node.Next)
	node.Next.Next = node
	node.Next = nil
	return newHead
}

/**
环型链表
 */
func detectCycle(head *Model.ListNode) *Model.ListNode {
	slow := head
	fast := head
	isCycle := false
	//先用快慢指针判断是否是环形
	for ; fast != nil && fast.Next != nil; {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			isCycle = true
			break
		}
	}

	if isCycle {
	    //慢指针不动，利用快指针找出环的大小
		cycleSize := 1
		fast = fast.Next
		for ; slow != fast; {
			fast = fast.Next
			cycleSize++
		}
        //根据环的大小，利用双指针找出环形入口，前后指针间隔为环的大小
		slow1 := head
		fast1 := head
		for ; cycleSize-1 > 0; {
			fast1 = fast1.Next
			cycleSize--
		}
        //找入口的关键判断条件
		for ; fast1.Next != slow1; {
			slow1 = slow1.Next;
			fast1 = fast1.Next;
		}
		return slow1
	}
	return nil
}

/**
合并K个链表
 */
func mergeKLists(lists []*ListNode) *ListNode {
if lists==nil{
return nil
}
   for i:=0;i<lists.size();i++{
   merge()
   }
}
