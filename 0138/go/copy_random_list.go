package leetcode

type Node struct {
    Val int
    Next *Node
    Random *Node
}

func copyRandomList(head *Node) *Node {
    index := 0
    currNode := head
    indexMap := make(map[*Node]int)
    for currNode != nil {
        indexMap[currNode] = index
        currNode = currNode.Next
        index++
    }

    newDummy := &Node{}
    prevNode, currNode := newDummy, head
    nodeSlice := []*Node{}
    indexSlice := []int{}
    for currNode != nil {
        newNode := &Node{
            Val: currNode.Val,
        }

        index := -1
        if currNode.Random != nil {
            index = indexMap[currNode.Random]
        }
        indexSlice = append(indexSlice, index)
        nodeSlice = append(nodeSlice, newNode)
        prevNode.Next = newNode
        prevNode = newNode

        currNode = currNode.Next
    }


    for i, n := range nodeSlice {
        index = indexSlice[i]
        if index != -1 {
            n.Random = nodeSlice[index]
        }
    }

    return newDummy.Next
}
