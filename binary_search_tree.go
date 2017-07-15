package binarysearchtree

const testVersion = 1

type SearchTreeData struct {
	left  *SearchTreeData
	data  int
	right *SearchTreeData
}

func Bst(d int) *SearchTreeData {
	std := new(SearchTreeData)
	std.data = d
	return std
}
func (std *SearchTreeData) Insert(d int) {
	childNode := Bst(d)
	currentNode := std
	for {
		if d <= currentNode.data {
			if currentNode.left == nil {
				currentNode.left = childNode
				return
			}
			currentNode = currentNode.left
		} else {
			if currentNode.right == nil {
				currentNode.right = childNode
				return
			}
			currentNode = currentNode.right
		}
	}
}
func (std *SearchTreeData) MapString(f func(int) string) (out []string) {
	outInts := printNode(std)
	out = make([]string, len(outInts))
	for i, outInt := range outInts {
		out[i] = f(outInt)
	}
	return
}
func (std *SearchTreeData) MapInt(f func(int) int) (out []int) {
	outInts := printNode(std)
	out = make([]int, len(outInts))
	for i, outInt := range outInts {
		out[i] = f(outInt)
	}
	return
}
func printNode(node *SearchTreeData) (out []int) {
	if node != nil {
		out = append(out, printNode(node.left)...)
		out = append(out, node.data)
		out = append(out, printNode(node.right)...)
	}
	return
}
