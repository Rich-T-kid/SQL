package DataStructures

import (
	"fmt"
	"math"
)

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}
type DiskTree[T comparable] interface {
	// n = # of nodes in tree , m
	//Adds a new element into the B+ tree and maintains the tree property.
	Insert(data T) //log n
	// Removes an element from the B+ tree. returns weather element was found or not
	Delete(data T) bool // log n
	// Searches for an element in the B+ tree using the key.
	Search(data T) bool // log n
	// Merges two nodes during deletion if necessary.
	Display() // BFS View of Tree O(n)
}

type BTree[T Ordered] struct {
	degree uint16
	root   *bNode[T]
	/*
		d = degree
		at most d children/ atleast math.floor(m/2) children
		at most d-1 keys/ atleast (math.floor(m/2) - 1) children
	*/
}
type bNode[T Ordered] struct {
	keys     []T         // Keys stored in this node, sorted
	children []*bNode[T] // Pointers to child nodes (empty for leaf nodes)
	parent   *bNode[T]   // Parent node
	leaf     bool        // Is this a leaf node?
	next     *bNode[T]   // Pointer to the next leaf node (for leaf nodes only)
}

func newBNode[T Ordered]() *bNode[T] {
	return &bNode[T]{
		keys:     make([]T, 0),
		children: make([]*bNode[T], 0),
		leaf:     true,
		next:     nil,
	}
}
func (b *BTree[T]) Search(item T) bool {
	if b.root == nil {
		return false
	}
	found, _, _ := b.searchNode(b.root, item)
	return found
}

/*
Iterates over all keys. if match is found return true, if target is greater than the current key and its not a leaf node, recusivly check the left subtre
if no key was greater than the target the key can only be in the right most sub tree. recusivly check right subtree
*/
func (b *BTree[T]) searchNode(node *bNode[T], target T) (bool, *bNode[T], int) {
	// currently this is iterative but we can assume that its sorted. Later optimize and include a binary search here
	for i, key := range node.keys {
		if key == target {
			return true, node, i
		} else if target < key {
			// belongs to left child of current key
			if node.leaf {
				return false, nil, 0
			}
			return b.searchNode(node.children[i], target)
		}
	}
	// if it isnt found in first m-1 sub children then its in the right most key
	if !node.leaf {
		return b.searchNode(node.children[len(node.children)-1], target)
	}
	return false, nil, 0

}

// Searches for element, if element exist nothing happends
// adds  item to key slice at correct node
// if over fill has happened at node it splits the node and reoganized the tree
func (b *BTree[T]) Insert(item T) {
	// Locate the leaf node and the position to insert
	exist, node, index := b.searchNode(b.root, item)
	if exist { // no need to add duplicates
		return
	}

	// Insert the key at the correct position
	//The first append copies keys before the insertion index.
	//The second append adds the new key and the remaining keys.
	node.keys = append(node.keys[:index], append([]T{item}, node.keys[index:]...)...)

	// Handle overfill (splitting)
	if node.overFill(b.degree) {
		b.Split(node)
	}
}

func (b *BTree[T]) Delete(item T) bool {

	exist, node, indexToRemove := b.searchNode(b.root, item)
	if !exist { // if it isnt found no need to do anything
		return false
	}
	// remove element from slice
	node.keys = append(node.keys[:indexToRemove], node.keys[indexToRemove+1:]...)
	//if undersized merge
	if node.underFill(b.degree) {
		b.Merge(node)
	}
	return true
}

func (b *BTree[T]) Split(node *bNode[T]) {
	mid := len(node.keys) / 2
	middleKey := node.keys[mid]

	// Create a new sibling node for the right half
	sibling := &bNode[T]{
		keys:     append([]T{}, node.keys[mid+1:]...),
		children: append([]*bNode[T]{}, node.children[mid+1:]...),
		leaf:     node.leaf,
		parent:   node.parent,
	}

	// Update the current node to keep the left half
	node.keys = node.keys[:mid]
	node.children = node.children[:mid+1]

	// If it's a leaf, adjust the linked list pointers
	if node.leaf {
		sibling.next = node.next
		node.next = sibling
	}

	// Handle propagation to the parent
	if node.parent == nil {
		// Splitting the root: Create a new root
		b.root = &bNode[T]{
			keys:     []T{middleKey},
			children: []*bNode[T]{node, sibling},
			leaf:     false,
		}
		node.parent = b.root
		sibling.parent = b.root
	} else {
		// Insert the middle key into the parent
		parent := node.parent
		insertIndex := 0
		for insertIndex < len(parent.keys) && parent.keys[insertIndex] < middleKey {
			insertIndex++
		}
		parent.keys = append(parent.keys[:insertIndex], append([]T{middleKey}, parent.keys[insertIndex:]...)...)
		parent.children = append(parent.children[:insertIndex+1], append([]*bNode[T]{sibling}, parent.children[insertIndex+1:]...)...)

		// Check for overfill in the parent
		if parent.overFill(b.degree) {
			b.Split(parent)
		}
	}
}

// Helper function for delete

func (b *BTree[T]) Merge(node *bNode[T]) {
	// Access the parent node directly
	parent := node.parent
	if parent == nil {
		return // No parent means no merging is required
	}

	// Find the index of the node in its parent's children
	var nodeIndex int
	for i, child := range parent.children {
		if child == node {
			nodeIndex = i
			break
		}
	}

	// Determine left or right sibling
	var sibling *bNode[T]
	var isLeftSibling bool
	if nodeIndex > 0 {
		// Prefer left sibling if it exists
		sibling = parent.children[nodeIndex-1]
		isLeftSibling = true
	} else if nodeIndex < len(parent.children)-1 {
		// Otherwise, use right sibling
		sibling = parent.children[nodeIndex+1]
		isLeftSibling = false
	}

	// Merge logic
	if sibling != nil {
		if isLeftSibling {
			// Merge node into left sibling
			sibling.keys = append(sibling.keys, parent.keys[nodeIndex-1])
			sibling.keys = append(sibling.keys, node.keys...)
			if !node.leaf {
				sibling.children = append(sibling.children, node.children...)
			}
			parent.keys = append(parent.keys[:nodeIndex-1], parent.keys[nodeIndex:]...)
			parent.children = append(parent.children[:nodeIndex], parent.children[nodeIndex+1:]...)
		} else {
			// Merge right sibling into node
			node.keys = append(node.keys, parent.keys[nodeIndex])
			node.keys = append(node.keys, sibling.keys...)
			if !node.leaf {
				node.children = append(node.children, sibling.children...)
			}
			parent.keys = append(parent.keys[:nodeIndex], parent.keys[nodeIndex+1:]...)
			parent.children = append(parent.children[:nodeIndex+1], parent.children[nodeIndex+2:]...)
		}
	}

	// Handle parent underfill
	if parent.underFill(b.degree) {
		if parent == b.root && len(parent.keys) == 0 {
			// Special case: root underflows and needs to shrink
			b.root = parent.children[0]
			b.root.parent = nil
		} else {
			b.Merge(parent) // Recursively handle parent merging
		}
	}
}

// Print out level order of B+ Tree
func (b *BTree[T]) Display() {
	if b.root == nil {
		fmt.Println("Empty Tree. No levels to Print")
		return
	}
	queue := NewQueue[*bNode[T]]()
	queue.Enqueue(b.root)
	levelMarker := NewQueue[*bNode[T]]() // To track level separation
	levelMarker.Enqueue(nil)             // Marker for level end

	var line string
	count := 0
	for !queue.IsEmpty() {
		currentNode := queue.Dequeue()
		if currentNode == nil {
			fmt.Printf("Level %d of B+ Tree: %s\n", count, line)
			line = ""
			count++

			// If there are more nodes in the queue, add another level marker
			if !queue.IsEmpty() {
				queue.Enqueue(nil)
			}
			continue
		}
		// Append keys of the current node
		for _, key := range currentNode.keys {
			line += fmt.Sprintf(" %v", key)
		}

		// Enqueue children of the current node
		for _, child := range currentNode.children {
			queue.Enqueue(child)
		}
	}
}

func (b *bNode[T]) overFill(degree uint16) bool {
	return uint16(len(b.keys)) >= degree
}
func (b *bNode[T]) underFill(degree uint16) bool {
	minKeys := uint16(math.Ceil(float64(degree)/2)) - 1
	return uint16(len(b.keys)) < minKeys
}

/*
degree T
all Nodes exepct root must have t-1 keys
root must contain a minimum of 1 key ROOT MAY NEVER BE EMPTY
all nodes can contain at most 2t - 1 keys

all keys must be in sorted order

insertions happend upward/ I.E insertions and deletions happend from the root
*/
