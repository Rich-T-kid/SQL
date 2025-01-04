package DataStructures

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	t.Run("Test AddFront", func(t *testing.T) {
		list := &LinkedList[int]{}
		list.Addfront(10) // Add to empty list
		list.Addfront(20) // Add to non-empty list
		list.Addfront(30) // Add another

		if list.head.value != 30 || list.tail.value != 10 {
			t.Errorf("AddFront failed. Expected head: 30, tail: 10. Got head: %d, tail: %d", list.head.value, list.tail.value)
		}
	})

	t.Run("Test Append", func(t *testing.T) {
		list := &LinkedList[int]{}
		list.Append(10) // Append to empty list
		list.Append(20) // Append to non-empty list
		list.Append(30) // Append another

		if list.head.value != 10 || list.tail.value != 30 {
			t.Errorf("Append failed. Expected head: 10, tail: 30. Got head: %d, tail: %d", list.head.value, list.tail.value)
		}
	})

	t.Run("Test PopHead", func(t *testing.T) {
		list := &LinkedList[int]{}
		list.Addfront(10)
		list.Addfront(20)

		node := list.PopHead()
		if node.value != 20 || list.head.value != 10 {
			t.Errorf("PopHead failed. Expected head: 10, popped: 20. Got head: %d, popped: %d", list.head.value, node.value)
		}

		node = list.PopHead()
		if node.value != 10 || list.head != nil {
			t.Errorf("PopHead failed. Expected head: nil, popped: 10. Got head: %v, popped: %d", list.head, node.value)
		}

		node = list.PopHead()
		if node != nil {
			t.Errorf("PopHead failed. Expected popped: nil. Got popped: %v", node)
		}
	})

	t.Run("Test PopTail", func(t *testing.T) {
		list := &LinkedList[int]{}
		list.Append(10)
		list.Append(20)

		node := list.PopTail()
		if node.value != 20 || list.tail.value != 10 {
			t.Errorf("PopTail failed. Expected tail: 10, popped: 20. Got tail: %d, popped: %d", list.tail.value, node.value)
		}

		node = list.PopTail()
		if node.value != 10 || list.tail != nil {
			t.Errorf("PopTail failed. Expected tail: nil, popped: 10. Got tail: %v, popped: %d", list.tail, node.value)
		}

		node = list.PopTail()
		if node != nil {
			t.Errorf("PopTail failed. Expected popped: nil. Got popped: %v", node)
		}
	})

	t.Run("Test Display", func(t *testing.T) {
		list := &LinkedList[int]{}
		for i := 1; i <= 15; i++ {
			list.Append(i)
		}
		list.Display() // Just ensuring no panic or error occurs
	})

	t.Run("Test Search", func(t *testing.T) {
		list := &LinkedList[int]{}
		list.Append(10)
		list.Append(20)
		list.Append(30)

		node := list.Search(20)
		if node == nil || node.value != 20 {
			t.Errorf("Search failed. Expected found node with value: 20. Got: %v", node)
		}

		node = list.Search(40)
		if node != nil {
			t.Errorf("Search failed. Expected nil for value: 40. Got: %v", node)
		}
	})

	t.Run("Test Clear", func(t *testing.T) {
		list := &LinkedList[int]{}
		list.Append(10)
		list.Append(20)
		list.Clear()

		if list.head != nil || list.tail != nil {
			t.Errorf("Clear failed. Expected head: nil, tail: nil. Got head: %v, tail: %v", list.head, list.tail)
		}
	})

	t.Run("Test Size", func(t *testing.T) {
		list := &LinkedList[int]{}
		if list.Size() != 0 {
			t.Errorf("Size failed. Expected 0 for empty list. Got: %d", list.Size())
		}

		list.Append(10)
		list.Append(20)
		if list.Size() != 2 {
			t.Errorf("Size failed. Expected 2. Got: %d", list.Size())
		}
	})

	t.Run("Test MergeSort", func(t *testing.T) {
		list := &LinkedList[int]{}
		list.Append(4)
		list.Append(2)
		list.Append(1)
		list.Append(3)

		list.head = list.MergeSort()
		sorted := list.ToSlice()

		expected := []int{1, 2, 3, 4}
		for i, v := range expected {
			if sorted[i] != v {
				t.Errorf("MergeSort failed. Expected: %v. Got: %v", expected, sorted)
				break
			}
		}
	})

	t.Run("Test Concate", func(t *testing.T) {
		list1 := &LinkedList[int]{}
		list1.Append(1)
		list1.Append(2)

		list2 := &LinkedList[int]{}
		list2.Append(3)
		list2.Append(4)

		result := list1.Concate(list2)
		expected := []int{1, 2, 3, 4}

		for i, v := range expected {
			if result[i] != v {
				t.Errorf("Concate failed. Expected: %v. Got: %v", expected, result)
				break
			}
		}
	})

	t.Run("Test Empty", func(t *testing.T) {
		list := &LinkedList[int]{}
		if !list.Empty() {
			t.Errorf("Empty failed. Expected true for empty list. Got: %v", list.Empty())
		}

		list.Append(10)
		if list.Empty() {
			t.Errorf("Empty failed. Expected false for non-empty list. Got: %v", list.Empty())
		}
	})
}

func TestAdditionalLinkedListCases(t *testing.T) {
	t.Run("AddFront - Adding Duplicate Values", func(t *testing.T) {
		list := &LinkedList[int]{}
		list.Addfront(10)
		list.Addfront(10) // Adding duplicate values
		list.Addfront(10)

		if list.Size() != 3 {
			t.Errorf("AddFront failed. Expected size: 3. Got: %d", list.Size())
		}

		if list.head.value != 10 || list.head.next.value != 10 || list.head.next.next.value != 10 {
			t.Errorf("AddFront failed. Expected all nodes to have value: 10.")
		}
	})

	t.Run("Append - Handling Different Data Types", func(t *testing.T) {
		strList := &LinkedList[string]{}
		strList.Append("Hello")
		strList.Append("World")

		if strList.Size() != 2 || strList.head.value != "Hello" || strList.tail.value != "World" {
			t.Errorf("Append failed for string list. Got head: %s, tail: %s", strList.head.value, strList.tail.value)
		}

		floatList := &LinkedList[float64]{}
		floatList.Append(3.14)
		floatList.Append(2.71)

		if floatList.Size() != 2 || floatList.head.value != 3.14 || floatList.tail.value != 2.71 {
			t.Errorf("Append failed for float list. Got head: %f, tail: %f", floatList.head.value, floatList.tail.value)
		}
	})

	t.Run("PopHead - Large List", func(t *testing.T) {
		list := &LinkedList[int]{}
		for i := 1; i <= 100; i++ {
			list.Append(i)
		}

		for i := 1; i <= 50; i++ {
			node := list.PopHead()
			if node.value != i {
				t.Errorf("PopHead failed at iteration %d. Expected: %d. Got: %d", i, i, node.value)
			}
		}

		if list.Size() != 50 {
			t.Errorf("PopHead failed. Expected size: 50 after 50 pops. Got: %d", list.Size())
		}
	})

	t.Run("PopTail - Empty List After Multiple Pops", func(t *testing.T) {
		list := &LinkedList[int]{}
		list.Append(1)
		list.Append(2)
		list.Append(3)

		list.PopTail()
		list.PopTail()
		list.PopTail()
		list.PopTail() // Additional pop on an already empty list

		if !list.Empty() {
			t.Errorf("PopTail failed. Expected empty list after popping all elements.")
		}
	})

	t.Run("Search - Non-Existent Elements", func(t *testing.T) {
		list := &LinkedList[int]{}
		list.Append(5)
		list.Append(10)
		list.Append(15)

		node := list.Search(20) // Non-existent value
		if node != nil {
			t.Errorf("Search failed. Expected nil for non-existent value. Got: %v", node)
		}
	})

	t.Run("MergeSort - Pre-Sorted List", func(t *testing.T) {
		list := &LinkedList[int]{}
		list.Append(1)
		list.Append(2)
		list.Append(3)
		list.Append(4)

		list.MergeSort()
		sorted := list.ToSlice()

		expected := []int{1, 2, 3, 4}
		for i, v := range expected {
			if sorted[i] != v {
				t.Errorf("MergeSort failed for pre-sorted list. Expected: %v. Got: %v", expected, sorted)
				break
			}
		}
	})

	t.Run("Merge - Unequal List Sizes", func(t *testing.T) {
		list1 := &LinkedList[int]{}
		list1.Append(1)
		list1.Append(3)
		list1.Append(5)

		list2 := &LinkedList[int]{}
		list2.Append(2)
		list2.Append(4)

		result := list1.Merge(list2)
		current := result
		expected := []int{1, 2, 3, 4, 5}

		for i, v := range expected {
			if current == nil || current.value != v {
				t.Errorf("Merge failed for unequal list sizes. Expected value: %d at index %d. Got: %v", v, i, current.value)
				break
			}
			current = current.next
		}
	})

	t.Run("Concate - Large Input Lists", func(t *testing.T) {
		list1 := &LinkedList[int]{}
		list2 := &LinkedList[int]{}

		for i := 1; i <= 1000; i++ {
			list1.Append(i)
		}
		for i := 1001; i <= 2000; i++ {
			list2.Append(i)
		}

		concatenated := list1.Concate(list2)
		if len(concatenated) != 2000 {
			t.Errorf("Concate failed for large input lists. Expected length: 2000. Got: %d", len(concatenated))
		}

		for i := 1; i <= 2000; i++ {
			if concatenated[i-1] != i {
				t.Errorf("Concate failed at index %d. Expected: %d. Got: %d", i-1, i, concatenated[i-1])
			}
		}
	})

	t.Run("Empty - After Clear", func(t *testing.T) {
		list := &LinkedList[int]{}
		list.Append(10)
		list.Append(20)
		list.Clear()

		if !list.Empty() {
			t.Errorf("Empty failed. Expected true after Clear. Got: %v", list.Empty())
		}
	})
}
func TestExtracases(t *testing.T) {
	t.Run("Test AddFront with Circular Reference", func(t *testing.T) {
		list := &LinkedList[int]{}
		list.Addfront(1)
		list.Addfront(2)
		list.Addfront(3)

		// Create a circular reference
		list.head.prev = list.tail
		list.tail.next = list.head

		// Check if AddFront can still handle it (safeguard)
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("AddFront caused a panic with circular reference: %v", r)
			}
		}()

		list.Addfront(4) // Should not break
	})

	t.Run("Test PopHead Repeatedly", func(t *testing.T) {
		list := &LinkedList[int]{}
		list.Append(1)
		list.Append(2)
		list.Append(3)

		// Pop elements one by one
		list.PopHead()
		list.PopHead()
		list.PopHead()

		if !list.Empty() {
			t.Errorf("PopHead repeatedly failed. Expected list to be empty, but it's not.")
		}

		// Try popping from an already empty list
		node := list.PopHead()
		if node != nil {
			t.Errorf("PopHead failed. Expected nil when popping from an empty list. Got: %v", node)
		}
	})

	t.Run("Test Concate Empty Lists", func(t *testing.T) {
		list1 := &LinkedList[int]{}
		list2 := &LinkedList[int]{}

		result := list1.Concate(list2)
		if len(result) != 0 {
			t.Errorf("Concate failed. Expected empty slice when concatenating two empty lists. Got: %v", result)
		}
	})
	t.Run("Test MergeSort on Already Sorted List", func(t *testing.T) {
		list := &LinkedList[int]{}
		list.Append(1)
		list.Append(2)
		list.Append(3)
		list.Append(4)

		list.MergeSort()
		sorted := list.ToSlice()

		expected := []int{1, 2, 3, 4}
		for i, v := range expected {
			if sorted[i] != v {
				t.Errorf("MergeSort failed on already sorted list. Expected: %v. Got: %v", expected, sorted)
				break
			}
		}
	})
	t.Run("Test Clear After Operations", func(t *testing.T) {
		list := &LinkedList[int]{}
		list.Append(10)
		list.Addfront(5)
		list.PopTail()
		list.Addfront(20)

		list.Clear()

		if !list.Empty() {
			t.Errorf("Clear failed after multiple operations. Expected list to be empty. Got: head=%v, tail=%v", list.head, list.tail)
		}
	})
	t.Run("Test Search for Non-Existent Element", func(t *testing.T) {
		list := &LinkedList[int]{}
		list.Append(1)
		list.Append(2)
		list.Append(3)

		node := list.Search(99) // Search for a value that doesn't exist
		if node != nil {
			t.Errorf("Search failed. Expected nil for non-existent value. Got: %v", node)
		}
	})
	t.Run("Test Merge Different Sized Lists", func(t *testing.T) {
		list1 := &LinkedList[int]{}
		list1.Append(1)
		list1.Append(3)

		list2 := &LinkedList[int]{}
		list2.Append(2)
		list2.Append(4)
		list2.Append(5)

		mergedHead := list1.Merge(list2)
		mergedList := &LinkedList[int]{head: mergedHead}

		result := mergedList.ToSlice()
		expected := []int{1, 2, 3, 4, 5}

		for i, v := range expected {
			if result[i] != v {
				t.Errorf("Merge failed with different sized lists. Expected: %v. Got: %v", expected, result)
				break
			}
		}
	})
	t.Run("Test Display with Single Element", func(t *testing.T) {
		list := &LinkedList[int]{}
		list.Append(42) // Single element

		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Display caused a panic with single element: %v", r)
			}
		}()

		list.Display() // Should print "node: 42" and "End Of Linked List"
	})
	t.Run("Test Size After Multiple Operations", func(t *testing.T) {
		list := &LinkedList[int]{}
		list.Append(1)
		list.Addfront(0)
		list.Append(2)
		list.PopTail()
		list.PopHead()

		size := list.Size()
		if size != 1 {
			t.Errorf("Size failed after multiple operations. Expected: 1. Got: %d", size)
		}
	})

}
