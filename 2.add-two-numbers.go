/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var lefts []*ListNode
	var pass, left int
	for {
		sum := l1.Val + l2.Val + pass
		pass = sum / 10
		left = sum % 10
		node := &ListNode{left, nil}
		lefts = append(lefts, node)
		if l1.Next == nil && l2.Next == nil {
			if pass != 0 {
				lefts = append(lefts, &ListNode{pass, nil})
			}
			break
		}

		if l1.Next == nil {
			l1 = new(ListNode)
		} else {
			l1 = l1.Next
		}
		if l2.Next == nil {
			l2 = new(ListNode)
		} else {
			l2 = l2.Next
		}
	}

	for i := 0; i < len(lefts)-1; i++ {
		lefts[i].Next = lefts[i+1]
	}
	return lefts[0]
}
