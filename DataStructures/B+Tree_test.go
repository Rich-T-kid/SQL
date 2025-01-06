package DataStructures

import (
	"testing"
)

// TestInsert validates insert functionality.
func TestInsert(t *testing.T) {
	btree := newBTree[int](4)

	// Insert elements and check tree structure.
	values := []int{10, 20, 5, 6, 15, 30, 25, 35}
	for _, v := range values {
		//fmt.Println("attempting to insert ", v)
		btree.Insert(v)
	}
	btree.Display()

	// Verify elements are present.
	for _, v := range values {
		if !btree.Search(v) {
			t.Errorf("Insert failed: %d not found in tree", v)
		}
	}

	// Test duplicate insertion.
	btree.Insert(10)
	count := countOccurrences(btree, 10)
	if count != 1 {
		t.Errorf("Duplicate insertion failed: expected 1, got %d", count)
	}
}
func TestEmptyTreeInsert(t *testing.T) {
	tree := newBTree[int](4) //into an empty tree
	tree.Insert(10)

	// Verify root contains the inserted key
	if len(tree.root.keys) != 1 || tree.root.keys[0] != 10 {
		t.Errorf("Insert into empty tree failed. Expected root keys [10], got %v", tree.root.keys)
	}
}

func TestEmptyTreeSearch(t *testing.T) {
	tree := newBTree[int](4)

	// n an empty tree
	if tree.Search(10) {
		t.Errorf("Search in empty tree failed. Expected false, got true")
	}
}

// TestDelete validates delete functionality.
func TestDelete(t *testing.T) {
	btree := &BTree[int]{degree: 3}

	// Insert elements.
	values := []int{10, 20, 5, 6, 15, 30, 25, 35}
	for _, v := range values {
		btree.Insert(v)
	}

	// Delete some elements.
	toDelete := []int{6, 15, 25}
	for _, v := range toDelete {
		if !btree.Delete(v) {
			t.Errorf("Delete failed: %d not deleted from tree", v)
		}
	}

	// Verify deleted elements are absent.
	for _, v := range toDelete {
		if btree.Search(v) {
			t.Errorf("Delete failed: %d still found in tree", v)
		}
	}

	// Verify remaining elements are still present.
	remaining := []int{10, 20, 5, 30, 35}
	for _, v := range remaining {
		if !btree.Search(v) {
			t.Errorf("Delete caused data loss: %d not found in tree", v)
		}
	}

	btree.Display()
}

// TestEdgeCases tests edge cases like empty trees and single-node trees.
func TestEdgeCases(t *testing.T) {
	btree := &BTree[int]{degree: 3}

	// Test search and delete on empty tree.
	if btree.Search(10) {
		t.Errorf("Search failed: found value in an empty tree")
	}
	if btree.Delete(10) {
		t.Errorf("Delete failed: deleted value from an empty tree")
	}

	// Insert a single element.
	btree.Insert(10)

	// Search for the single element.
	if !btree.Search(10) {
		t.Errorf("Search failed: 10 not found in tree")
	}

	// Delete the single element.
	if !btree.Delete(10) {
		t.Errorf("Delete failed: 10 not deleted from tree")
	}

	// Verify the tree is empty again.
	if btree.Search(10) {
		t.Errorf("Search failed: 10 found in empty tree")
	}
}

func TestDisplay(t *testing.T) {
	btree := &BTree[int]{degree: 4}

	// Insert values.
	values := []int{50, 20, 70, 10, 30, 60, 80, 90, 40}
	for _, v := range values {
		btree.Insert(v)
	}

	// Display the tree.
	btree.Display()
}

// TestRandomized validates random insertions and deletions.

// TestDisplay ensures the display function works without errors.
// Helper function: count occurrences of a value in the tree.
func countOccurrences(btree *BTree[int], value int) int {
	count := 0
	queue := NewQueue[*bNode[int]]()
	queue.Enqueue(btree.root)

	for !queue.IsEmpty() {
		node := queue.Dequeue()
		if node == nil {
			continue
		}

		for _, key := range node.keys {
			if key == value {
				count++
			}
		}

		for _, child := range node.children {
			queue.Enqueue(child)
		}
	}

	return count
}
