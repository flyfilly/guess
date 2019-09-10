package guess

// Node is used to construct a BST
type Node struct {
	left  *Node
	value int
	right *Node
}

func (node *Node) search(data int) bool {
	if data == node.value {
		return true
	} else if data <= node.value && node.left != nil {
		return node.left.search(data)
	} else if node.right != nil {
		return node.right.search(data)
	}
	return false
}

func (node *Node) insert(data int) {
	if node == nil {
		return
	} else if data <= node.value {
		if node.left == nil {
			node.left = &Node{value: data, left: nil, right: nil}
		} else {
			node.left.insert(data)
		}
	} else {
		if node.right == nil {
			node.right = &Node{value: data, left: nil, right: nil}
		} else {
			node.right.insert(data)
		}
	}
}

// Play determines strong and weak matches between 2 integer arrays
func Play(code, guess []int) (int, int) {
	strong, weak := 0, 0
	weakCandidates := []int{}
	var tree Node

	for i := range code {
		if code[i] == guess[i] {
			strong++
		} else {
			weakCandidates = append(weakCandidates, guess[i])
			tree.insert(code[i])
		}
	}

	for _, candidate := range weakCandidates {
		if tree.search(candidate) {
			weak++
		}
	}

	return strong, weak
}
