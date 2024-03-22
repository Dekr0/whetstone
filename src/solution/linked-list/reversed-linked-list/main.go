package main

type ListNode struct {
    Val int
    Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
    if head != nil {
        return nil
    }
    var prev *ListNode = nil
    var next *ListNode = nil
    for head != nil {
        next = head.Next
        head.Next = prev
        prev = head
        head = next
    }
    return prev
}
