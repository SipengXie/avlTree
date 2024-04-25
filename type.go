package avltree

// 维护st
type Node struct {
	payload uint64

	lson, rson *Node

	height int
}

func (n *Node) getHeight() int {
	if n == nil {
		return 0
	}
	return n.height
}

func (n *Node) recalculateHeight() {
	n.height = 1 + max(n.lson.getHeight(), n.rson.getHeight())
}

// Checks if node is balanced and rebalance
func (n *Node) rebalanceTree() *Node {
	if n == nil {
		return n
	}
	n.recalculateHeight()

	// check balance factor and rotatelson if rson-heavy and rotaterson if lson-heavy
	balanceFactor := n.lson.getHeight() - n.rson.getHeight()
	if balanceFactor == -2 {
		// check if child is lson-heavy and rotaterson first
		if n.rson.lson.getHeight() > n.rson.rson.getHeight() {
			n.rson = n.rson.rotaterson()
		}
		return n.rotatelson()
	} else if balanceFactor == 2 {
		// check if child is rson-heavy and rotatelson first
		if n.lson.rson.getHeight() > n.lson.lson.getHeight() {
			n.lson = n.lson.rotatelson()
		}
		return n.rotaterson()
	}
	return n
}

// Rotate nodes lson to balance node
func (n *Node) rotatelson() *Node {
	newRoot := n.rson
	n.rson = newRoot.lson
	newRoot.lson = n
	n.recalculateHeight()
	newRoot.recalculateHeight()
	return newRoot
}

// Rotate nodes rson to balance node
func (n *Node) rotaterson() *Node {
	newRoot := n.lson
	n.lson = newRoot.rson
	newRoot.rson = n
	n.recalculateHeight()
	newRoot.recalculateHeight()
	return newRoot
}

func (n *Node) add(st uint64) *Node {
	if n == nil {
		return &Node{payload: st, height: 1}
	}
	if st < n.payload {
		n.lson = n.lson.add(st)
	} else {
		n.rson = n.rson.add(st)
	}
	return n.rebalanceTree()
}

func (n *Node) remove(st uint64) *Node {
	if n == nil {
		return nil
	}
	if st < n.payload {
		n.lson = n.lson.remove(st)
	} else if st > n.payload {
		n.rson = n.rson.remove(st)
	} else {
		if n.lson != nil && n.rson != nil {
			minNode := n.rson.findMin()
			n.payload = minNode.payload
			n.rson = n.rson.remove(minNode.payload)
		} else if n.lson != nil {
			n = n.lson
		} else if n.rson != nil {
			n = n.rson
		} else {
			n = nil
			return n
		}
	}
	return n.rebalanceTree()
}

func (n *Node) findMin() *Node {
	if n.lson == nil {
		return n
	}
	return n.lson.findMin()
}

// 寻找小于等于st的最大节点
func (n *Node) findMaxLessThan(st uint64) *Node {
	if n == nil {
		return nil
	}
	if n.payload == st {
		return n
	} else if n.payload < st {
		if n.rson == nil {
			return n
		}
		// n是一个可能的答案，我们还要去右子树找
		// 若右子树找不到答案，则返回n
		tmp := n.rson.findMaxLessThan(st)
		if tmp == nil {
			return n
		}
		return tmp
	} else {
		// n.payload > st
		// lson == nil 会返回nil
		return n.lson.findMaxLessThan(st)
	}
}

type AVLTree struct {
	root *Node
}

func NewTree() *AVLTree {
	return &AVLTree{root: nil}
}

func (t *AVLTree) Add(st uint64) {
	t.root = t.root.add(st)
}

func (t *AVLTree) Remove(st uint64) {
	t.root = t.root.remove(st)
}

func (t *AVLTree) FindMaxLessThan(st uint64) *Node {
	return t.root.findMaxLessThan(st)
}
