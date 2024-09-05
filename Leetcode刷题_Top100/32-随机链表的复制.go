package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

var cachedNode map[*Node]*Node

func copyRandomList(head *Node) *Node {
	cachedNode = map[*Node]*Node{}
	return deepCopy(head)

}

// 对每一个节点进行深度拷贝
func deepCopy(node *Node) *Node {
	if node == nil {
		return nil
	}
	if n, has := cachedNode[node]; has {
		return n
	}
	newNode := &Node{Val: node.Val}        //对传入节点的值进行复制
	cachedNode[node] = newNode             //当前传入节点已经拷贝过，加入map
	newNode.Next = deepCopy(node.Next)     //对当前传入节点的下一个node进行拷贝
	newNode.Random = deepCopy(node.Random) //同理
	return newNode
}
