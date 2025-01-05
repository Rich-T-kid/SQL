package DataStructures

import "testing"

// Test with integers
func TestQueueWithIntegers(t *testing.T) {
	testQueueInt := NewQueue[int]()
	if !testQueueInt.IsEmpty() {
		t.Errorf("Expected queue to be empty, got %v", testQueueInt.IsEmpty())
	}

	testQueueInt.Enqueue(10)
	testQueueInt.Enqueue(20)
	testQueueInt.Enqueue(30)
	if testQueueInt.Size() != 3 {
		t.Errorf("Expected queue size to be 3, got %d", testQueueInt.Size())
	}
	if testQueueInt.Peek() != 10 {
		t.Errorf("Expected queue peek to be 10, got %d", testQueueInt.Peek())
	}

	// Proper Dequeue sequence
	for !testQueueInt.IsEmpty() {
		dequeued := testQueueInt.Dequeue()
		t.Logf("Dequeued: %d", dequeued)
	}
	if !testQueueInt.IsEmpty() {
		t.Errorf("Expected queue to be empty after dequeueing all elements, got %v", testQueueInt.IsEmpty())
	}

	// Edge Case: Dequeue on empty queue
	if testQueueInt.IsEmpty() {
		t.Log("Queue is empty. Cannot dequeue.")
	}
}

// Test with strings
func TestQueueWithStrings(t *testing.T) {
	testQueueStr := NewQueue[string]()
	if !testQueueStr.IsEmpty() {
		t.Errorf("Expected queue to be empty, got %v", testQueueStr.IsEmpty())
	}

	testQueueStr.Enqueue("first")
	testQueueStr.Enqueue("second")
	testQueueStr.Enqueue("third")
	if testQueueStr.Size() != 3 {
		t.Errorf("Expected queue size to be 3, got %d", testQueueStr.Size())
	}
	if testQueueStr.Peek() != "first" {
		t.Errorf("Expected queue peek to be 'first', got %s", testQueueStr.Peek())
	}

	// Proper Dequeue sequence
	for !testQueueStr.IsEmpty() {
		dequeued := testQueueStr.Dequeue()
		t.Logf("Dequeued: %s", dequeued)
	}
	if !testQueueStr.IsEmpty() {
		t.Errorf("Expected queue to be empty after dequeueing all elements, got %v", testQueueStr.IsEmpty())
	}

	// Edge Case: Dequeue on empty queue
	if testQueueStr.IsEmpty() {
		t.Log("Queue is empty. Cannot dequeue.")
	}
}

// Test with custom struct
type Point struct {
	x, y int
}

func TestQueueWithCustomStruct(t *testing.T) {
	testQueueStruct := NewQueue[Point]()
	if !testQueueStruct.IsEmpty() {
		t.Errorf("Expected queue to be empty, got %v", testQueueStruct.IsEmpty())
	}

	testQueueStruct.Enqueue(Point{1, 2})
	testQueueStruct.Enqueue(Point{3, 4})
	testQueueStruct.Enqueue(Point{5, 6})
	if testQueueStruct.Size() != 3 {
		t.Errorf("Expected queue size to be 3, got %d", testQueueStruct.Size())
	}
	if testQueueStruct.Peek() != (Point{1, 2}) {
		t.Errorf("Expected queue peek to be {1, 2}, got %+v", testQueueStruct.Peek())
	}

	// Proper Dequeue sequence
	for !testQueueStruct.IsEmpty() {
		dequeued := testQueueStruct.Dequeue()
		t.Logf("Dequeued: %+v", dequeued)
	}
	if !testQueueStruct.IsEmpty() {
		t.Errorf("Expected queue to be empty after dequeueing all elements, got %v", testQueueStruct.IsEmpty())
	}
}

// Test edge cases
func TestQueueEdgeCases(t *testing.T) {
	largeQueue := NewQueue[int]()
	for i := 0; i < 255; i++ {
		largeQueue.Enqueue(i)
	}
	expected := 255
	if largeQueue.Size() != uint8(expected) {
		t.Errorf("Expected queue size to be %d, got %d", expected, largeQueue.Size())
	}
	for !largeQueue.IsEmpty() {
		largeQueue.Dequeue()
	}
	if !largeQueue.IsEmpty() {
		t.Errorf("Expected queue to be empty after dequeueing all elements, got %v", largeQueue.IsEmpty())
	}

	// Test IsEmpty on a non-empty queue
	largeQueue.Enqueue(100)
	if largeQueue.IsEmpty() {
		t.Errorf("Expected queue to be non-empty, got %v", largeQueue.IsEmpty())
	}
}

func TestDequeue(t *testing.T) {
	// Test with integers
	intDequeue := NewDequeue[int]()

	// Initial state
	if !intDequeue.IsEmpty() {
		t.Errorf("Expected dequeue to be empty initially")
	}

	// Enqueue elements
	intDequeue.Enqueue(1)
	intDequeue.Enqueue(2)
	intDequeue.Enqueue(3)

	if intDequeue.Size() != 3 {
		t.Errorf("Expected dequeue size to be 3, got %d", intDequeue.Size())
	}

	// Peek front
	if intDequeue.Peek() != 1 {
		t.Errorf("Expected front to be 1, got %d", intDequeue.Peek())
	}

	// Insert at rear
	intDequeue.InsertRear(0)
	if intDequeue.Peek() != 0 {
		t.Errorf("Expected front to be 0 after InsertRear, got %d", intDequeue.Peek())
	}

	// Dequeue front
	if intDequeue.Dequeue() != 0 {
		t.Errorf("Expected Dequeue to return 0")
	}

	if intDequeue.Size() != 3 {
		t.Errorf("Expected dequeue size to be 3 after one dequeue, got %d", intDequeue.Size())
	}

	// Pop rear
	if intDequeue.PopRear() != 3 {
		t.Errorf("Expected PopRear to return 3")
	}

	if intDequeue.Size() != 2 {
		t.Errorf("Expected dequeue size to be 2 after one PopRear, got %d", intDequeue.Size())
	}

	// Final state
	intDequeue.Dequeue()
	intDequeue.Dequeue()

	if !intDequeue.IsEmpty() {
		t.Errorf("Expected dequeue to be empty after removing all elements")
	}

	// Edge cases
	t.Run("Dequeue on empty", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected panic when dequeueing from an empty dequeue")
			}
		}()
		intDequeue.Dequeue()
	})

	t.Run("PopRear on empty", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected panic when popping rear from an empty dequeue")
			}
		}()
		intDequeue.PopRear()
	})

	// Test with strings
	strDequeue := NewDequeue[string]()
	strDequeue.Enqueue("A")
	strDequeue.Enqueue("B")
	strDequeue.InsertRear("Z")

	if strDequeue.Dequeue() != "Z" {
		t.Errorf("Expected Dequeue to return 'Z'")
	}
	if strDequeue.PopRear() != "B" {
		t.Errorf("Expected PopRear to return 'B'")
	}
	if strDequeue.Dequeue() != "A" {
		t.Errorf("Expected Dequeue to return 'A'")
	}

	// Large test case
	largeDequeue := NewDequeue[int]()
	for i := 0; i < 1000; i++ {
		largeDequeue.Enqueue(i)
	}
	for i := 999; i >= 0; i-- {
		if largeDequeue.PopRear() != i {
			t.Errorf("Expected PopRear to return %d", i)
		}
	}
}
