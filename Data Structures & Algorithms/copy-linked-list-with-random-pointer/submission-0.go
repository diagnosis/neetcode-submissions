/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node{
	hMap := make(map[*Node]*Node)
	curr := head
	for curr != nil{
		hMap[curr] = &Node{Val: curr.Val}
		curr = curr.Next
	}
	curr = head
	for curr != nil {
		hMap[curr].Next = hMap[curr.Next]
		hMap[curr].Random = hMap[curr.Random]
		curr = curr.Next
	}
	return hMap[head]
}
