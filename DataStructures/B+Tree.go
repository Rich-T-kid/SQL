package DataStructures

import (
	"fmt"
	"math"
)

var (
	defualtDegree = 4
)

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

var defaultDegree = 4

// DiskTree is just an interface
type DiskTree[T comparable] interface {
	Insert(data T)
	Delete(data T) bool
	Search(data T) bool
	Display()
}

// BTree is our B+ Tree structure
type BTree[T Ordered] struct {
	degree uint16
	root   *bNode[T]
}

// bNode represents a node in the B+ Tree
type bNode[T Ordered] struct {
	keys     []T         // sorted keys
	children []*bNode[T] // child pointers (empty if leaf)
	parent   *bNode[T]   // parent pointer
	leaf     bool        // is leaf node?
	next     *bNode[T]   // linked list pointer for leaves
}

// newBNode creates an empty leaf node
func newBNode[T Ordered]() *bNode[T] {
	return &bNode[T]{
		keys:     make([]T, 0),
		children: make([]*bNode[T], 0),
		leaf:     true,
		next:     nil,
	}
}

// newBTree creates a new B+ Tree of the given degree
func newBTree[T Ordered](degree uint16) *BTree[T] {
	// Start with an empty leaf as root
	root := newBNode[T]()
	root.leaf = true
	return &BTree[T]{
		degree: degree,
		root:   root,
	}
}

// Search checks if 'item' exists in the tree
func (b *BTree[T]) Search(item T) bool {
	if b.root == nil {
		// Should not happen if we always init root, but safe check
		return false
	}
	found, _, _ := b.searchNode(b.root, item)
	return found
}

// searchNode returns:
//
//	(found=true, node, index) if 'item' exists at node.keys[index]
//	(found=false, node, index) if 'item' does not exist, but should be inserted at index in node.keys
func (b *BTree[T]) searchNode(node *bNode[T], target T) (bool, *bNode[T], int) {
	if node == nil {
		return false, nil, 0
	}

	for i, key := range node.keys {
		if key == target {
			// Found exact match
			return true, node, i
		}
		if target < key {
			// Should go to the left child if not leaf
			if node.leaf {
				return false, node, i
			}
			return b.searchNode(node.children[i], target)
		}
	}
	// If target > all keys in node
	if node.leaf {
		return false, node, len(node.keys)
	}
	return b.searchNode(node.children[len(node.children)-1], target)
}

// Insert adds a new 'item' to the B+ Tree
func (b *BTree[T]) Insert(item T) {
	// If tree is uninitialized (edge case)
	if b.root == nil {
		b.root = newBNode[T]()
		b.root.keys = append(b.root.keys, item)
		return
	}

	// 1. Find the correct position (leaf node + index)
	exist, node, index := b.searchNode(b.root, item)
	if exist {
		// No duplicates
		return
	}
	if node == nil {
		panic("searchNode returned nil node on Insert")
	}
	// 2. Insert the key in the leaf at 'index'
	node.keys = append(node.keys[:index], append([]T{item}, node.keys[index:]...)...)

	// 3. Check for overfill => split
	if node.overFill(b.degree) {
		b.Split(node)
	}
}

// Delete removes 'item' from the tree, returning true if found/deleted
// Public Delete: wraps our internal recursive method.
func (b *BTree[T]) Delete(item T) bool {
	if b.root == nil {
		return false // Empty tree
	}

	deleted := b.deleteKey(b.root, item)
	if !deleted {
		return false // Key not found
	}

	// If the root became empty and has children, shrink the tree height
	if len(b.root.keys) == 0 && len(b.root.children) > 0 {
		b.root = b.root.children[0]
		b.root.parent = nil
	}
	return true
}

// Recursively delete 'item' from 'node' or its children.
func (b *BTree[T]) deleteKey(node *bNode[T], item T) bool {
	// 1. Find the position 'i' of 'item' or where 'item' should be in node.keys
	i := 0
	for i < len(node.keys) && node.keys[i] < item {
		i++
	}

	// 2. If 'item' found in node.keys[i]
	if i < len(node.keys) && node.keys[i] == item {
		// Case A: node is a LEAF
		if node.leaf {
			// Remove the key from the leaf
			node.keys = append(node.keys[:i], node.keys[i+1:]...)
		} else {
			// Case B: node is INTERNAL
			// We must swap 'item' with the predecessor or successor key
			// and then delete that key from the child leaf.

			// We'll pick the predecessor for demonstration:
			predNode := b.getPredecessorNode(node.children[i]) // rightmost leaf in left subtree
			predKey := predNode.keys[len(predNode.keys)-1]

			// Swap
			node.keys[i] = predKey
			// Recursively delete 'predKey' from predNode (which is a leaf)
			b.deleteKey(predNode, predKey)
		}
	} else {
		// Not found in this node
		// If leaf, it's not in the tree
		if node.leaf {
			return false
		}
		// Otherwise, descend into children[i]
		b.deleteKey(node.children[i], item)
	}

	// After removal, check if node underflows
	if node != nil && node != b.root && node.underFill(b.degree) {
		b.Merge(node)
	}
	return true
}

// getPredecessorNode traverses to the rightmost leaf in 'curNode'
func (b *BTree[T]) getPredecessorNode(curNode *bNode[T]) *bNode[T] {
	// We go down curNode's last child pointer repeatedly
	// until we reach a leaf. That leaf's last key is the predecessor.
	for !curNode.leaf {
		curNode = curNode.children[len(curNode.children)-1]
	}
	return curNode
}

// overFill checks if node has >= 'degree' keys
func (n *bNode[T]) overFill(degree uint16) bool {
	return len(n.keys) >= int(degree)
}

// underFill checks if node has fewer than floor(degree/2) keys
// e.g., for degree=4, node must have >=2 keys. So if <2 => underfill
func (n *bNode[T]) underFill(degree uint16) bool {
	minKeys := int(math.Ceil(float64(degree)/2.0)) - 1
	// e.g., degree=4 => minKeys=1
	return len(n.keys) < minKeys
}

// Split handles overfilled nodes
func (b *BTree[T]) Split(node *bNode[T]) {
	mid := len(node.keys) / 2
	middleKey := node.keys[mid]

	sibling := &bNode[T]{
		keys:   append([]T{}, node.keys[mid+1:]...),
		leaf:   node.leaf,
		parent: node.parent,
	}

	// Left node keeps left half
	node.keys = node.keys[:mid]

	if !node.leaf {
		// Split children for internal node
		sibling.children = append([]*bNode[T]{}, node.children[mid+1:]...)
		node.children = node.children[:mid+1]

		// Reassign parents
		for _, child := range sibling.children {
			child.parent = sibling
		}
	} else {
		// Leaf node: fix 'next' pointer
		sibling.next = node.next
		node.next = sibling
	}

	if node.parent == nil {
		// Splitting root
		newRoot := &bNode[T]{
			keys:     []T{middleKey},
			leaf:     false,
			children: []*bNode[T]{node, sibling},
		}
		node.parent = newRoot
		sibling.parent = newRoot
		b.root = newRoot
	} else {
		// Insert 'middleKey' into parent
		parent := node.parent

		insertPos := 0
		for insertPos < len(parent.keys) && parent.keys[insertPos] < middleKey {
			insertPos++
		}
		parent.keys = append(parent.keys[:insertPos],
			append([]T{middleKey}, parent.keys[insertPos:]...)...,
		)

		parent.children = append(
			parent.children[:insertPos+1],
			append([]*bNode[T]{sibling}, parent.children[insertPos+1:]...)...,
		)

		if parent.overFill(b.degree) {
			b.Split(parent)
		}
	}
}

// Merge handles underfilled nodes
func (b *BTree[T]) Merge(node *bNode[T]) {
	parent := node.parent
	if parent == nil {
		// If node is root, no merge needed
		return
	}

	nodeIndex := findChildIndex(parent, node)
	// If we can merge with left sibling, do so; else merge with right sibling
	if nodeIndex > 0 {
		leftSibling := parent.children[nodeIndex-1]
		mergeRightIntoLeft(leftSibling, node, parent, nodeIndex-1, b)
	} else {
		rightSibling := parent.children[nodeIndex+1]
		mergeRightIntoLeft(node, rightSibling, parent, nodeIndex, b)
	}
}

// findChildIndex locates 'child' in 'parent.children'
func findChildIndex[T Ordered](parent *bNode[T], child *bNode[T]) int {
	for i, c := range parent.children {
		if c == child {
			return i
		}
	}
	return -1
}

// mergeRightIntoLeft merges 'right' node into 'left' node
// then removes 'right' from the parent
func mergeRightIntoLeft[T Ordered](
	left *bNode[T],
	right *bNode[T],
	parent *bNode[T],
	parentKeyIndex int,
	b *BTree[T],
) {
	// Key in parent that separates left and right
	separatingKey := parent.keys[parentKeyIndex]

	// If not leaf, copy separatingKey into left
	if !left.leaf {
		left.keys = append(left.keys, separatingKey)
	}
	// Merge right.keys into left
	left.keys = append(left.keys, right.keys...)

	if !left.leaf {
		// Merge children
		left.children = append(left.children, right.children...)
		for _, child := range right.children {
			child.parent = left
		}
	} else {
		// If leaf, fix the leaf chain
		left.next = right.next
	}

	// Remove 'separatingKey' from parent
	parent.keys = append(
		parent.keys[:parentKeyIndex],
		parent.keys[parentKeyIndex+1:]...,
	)

	// Remove 'right' pointer from parent.children
	rightIndex := findChildIndex(parent, right)
	parent.children = append(
		parent.children[:rightIndex],
		parent.children[rightIndex+1:]...,
	)

	// Check if parent is underfilled
	if parent.underFill(b.degree) {
		if parent == b.root && len(parent.keys) == 0 {
			// If parent is root and empty => shrink
			b.root = parent.children[0]
			b.root.parent = nil
		} else {
			b.Merge(parent)
		}
	}
}

// Display prints the tree in a level-order (BFS) format
func (b *BTree[T]) Display() {
	if b.root == nil {
		fmt.Println("Empty tree. No levels to print.")
		return
	}
	queue := NewQueue[*bNode[T]]()
	queue.Enqueue(b.root)
	queue.Enqueue(nil)

	level := 0
	var line []string
	for !queue.IsEmpty() {
		node := queue.Dequeue()
		if node == nil {
			fmt.Printf("Level %d: %s\n", level, join(line, ", "))
			line = []string{}
			level++
			if !queue.IsEmpty() {
				queue.Enqueue(nil)
			}
			continue
		}
		// Collect keys for this level
		for _, key := range node.keys {
			line = append(line, fmt.Sprintf("%v", key))
		}
		// Enqueue children
		for _, child := range node.children {
			queue.Enqueue(child)
		}
	}
}

// join is a simple utility to join string slices with a separator
func join(elements []string, sep string) string {
	switch len(elements) {
	case 0:
		return ""
	case 1:
		return elements[0]
	}
	out := elements[0]
	for i := 1; i < len(elements); i++ {
		out += sep + elements[i]
	}
	return out
}
