package pointer

type ListNode struct {
	Val  int
	Next *ListNode
}

func GenList(nums []int) *ListNode {
	var head, cur *ListNode
	for index, num := range nums {
		if index == 0 {
			head = &ListNode{num, nil}
			cur = head
		} else {
			cur.Next = &ListNode{num, nil}
			cur = cur.Next
		}
	}

	return head
}

//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//
//}

func GetLastVal(l *ListNode) int {
	cur := l
	for i := 1; i > 0; i++ {
		if cur.Next == nil {
			break
		}
		cur = cur.Next
	}
	return cur.Val
}
