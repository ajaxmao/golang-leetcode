/*
 * [21] Merge Two Sorted Lists
 *
 * https://leetcode.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (41.63%)
 * Total Accepted:    353.2K
 * Total Submissions: 848.1K
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * Merge two sorted linked lists and return it as a new list. The new list
 * should be made by splicing together the nodes of the first two lists.
 *
 * Example:
 *
 * Input: 1->2->4, 1->3->4
 * Output: 1->1->2->3->4->4
 *
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	L := []*ListNode{}
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			L = append(L, l1)
			l1 = l1.Next
		} else {
			L = append(L, l2)
			l2 = l2.Next
		}
	}
	for i := 0; i < len(L)-1; i++ {
		L[i].Next = L[i+1]
	}
	if l1 == nil {
		L[len(L)-1].Next = l2
	}
	if l2 == nil {
		L[len(L)-1].Next = l1
	}

	return L[0]
}
