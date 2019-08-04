package main

import (
	"container/list"
	"fmt"
	"goleetcode/src/main/Model"
	"strings"
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
			slow1 = slow1.Next
			fast1 = fast1.Next
		}
		return slow1
	}
	return nil
}

/**
合并K个链表
*/
func mergeKLists(lists []*Model.ListNode) *Model.ListNode {
	if lists == nil {
		return nil
	}
	var mergeListNode *Model.ListNode
	for i := 0; i < len(lists); i++ {
		mergeListNode = merge(mergeListNode, lists[i])
	}
	return mergeListNode
}

/**
两数之和
*/
func twoSum(nums []int, target int) []int {
	mapObj := map[int]int{}
	for i := range nums {
		x := target - nums[i]
		n, ok := mapObj[x]
		if ok {
			return []int{n, i}
		} else {
			mapObj[nums[i]] = i
		}
	}
	return nil
}

/**
有效的括号
*/
func isValid(s string) bool {
	m := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}

	l := list.New()
	for i := 0; i < len(s); i++ {
		msg := s[i : i+1]
		if msg == "(" || msg == "[" || msg == "{" {
			l.PushBack(msg)
		}

		if msg == ")" || msg == "]" || msg == "}" {
			if l.Len() == 0 {
				return false
			}
			s := l.Back()
			if s.Value == m[msg] {
				l.Remove(s)
			} else {
				return false
			}
		}
	}

	if l.Len() == 0 {
		return true
	} else {
		return true
	}
}

/**
合并两个有序链表
*/
func mergeTwoLists(l1 *Model.ListNode, l2 *Model.ListNode) *Model.ListNode {
	result := &Model.ListNode{Val: 0}
	result2 := result
	for ; l1 != nil || l2 != nil; {
		if l1 == nil {
			result2.Next = l2
			l2 = l2.Next
		} else if l2 == nil {
			result2.Next = l1
			l1 = l1.Next
		} else if l1.Val < l2.Val {
			result2.Next = l1
			l1 = l1.Next
		} else {
			result2.Next = l2
			l2 = l2.Next
		}
		result2 = result2.Next
	}
	return result.Next
}

/**
单词规律
*/
func wordPattern(pattern string, str string) bool {
	array := strings.Fields(str)
	if len(array) != len(pattern) {
		return false
	}
	hash := make(map[byte]string)
	hash2 := make(map[string]byte)
	for i := 0; i < len(pattern); i++ {
		v, ok := hash[pattern[i]]
		v2, ok2 := hash2[array[i]]
		if ok && v != array[i] || ok2 && v2 != pattern[i] {
			return false
		} else {
			hash[pattern[i]] = array[i]
			hash2[array[i]] = pattern[i]
		}
	}
	return true
}

/**
移动零
*/
func moveZeroes(nums []int) {
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == 0 {
			for j := i; j < len(nums)-1 && nums[j+1] != 0; j++ {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}
