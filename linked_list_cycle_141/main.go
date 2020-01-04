package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// To represent a cycle in the given linked list, we use an integer pos which represents the position (0-indexed)
// (in the linked list where tail connects to. If pos is -1, then there is no cycle in the linked list.)
func hasCycle(head *ListNode) bool {
	// 現在のノードのポインタがnil 次がnil だったらループしていない
	if head == nil || head.Next == nil {
		return false
	}

	slow := head      // 現在のポインタ
	fast := head.Next // 次のノードのポインタ

	for slow != fast {
		// fastがnil fastの次がnilだったらループしていない
		if fast == nil || fast.Next == nil {
			return false
		}

		// 進み方に差をつけて比較することでループしているといずれ同じノードを比較する
		slow = slow.Next      // slowは遅れて進む
		fast = fast.Next.Next // fastがどんどん先に進んでいく
	}

	return true
}
