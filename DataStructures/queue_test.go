package DataStructures

import (
	"fmt"
)

func main() {
	// Test with integers
	fmt.Println("Testing Queue with integers...")
	testQueueInt := NewQueue[int]()
	fmt.Printf("Queue empty? %v\n", testQueueInt.IsEmpty())
	testQueueInt.Enqueue(10)
	testQueueInt.Enqueue(20)
	testQueueInt.Enqueue(30)
	fmt.Printf("Queue size after enqueue: %d\n", testQueueInt.Size())
	fmt.Printf("Queue peek: %d\n", testQueueInt.Peek())

	// Proper Dequeue sequence
	for !testQueueInt.IsEmpty() {
		fmt.Printf("Dequeued: %d\n", testQueueInt.Dequeue())
	}
	fmt.Printf("Queue empty after dequeueing all elements? %v\n", testQueueInt.IsEmpty())

	// Edge Case: Dequeue on empty queue
	if testQueueInt.IsEmpty() {
		fmt.Println("Queue is empty. Cannot dequeue.")
	}

	// Test with strings
	fmt.Println("\nTesting Queue with strings...")
	testQueueStr := NewQueue[string]()
	fmt.Printf("Queue empty? %v\n", testQueueStr.IsEmpty())
	testQueueStr.Enqueue("first")
	testQueueStr.Enqueue("second")
	testQueueStr.Enqueue("third")
	fmt.Printf("Queue size after enqueue: %d\n", testQueueStr.Size())
	fmt.Printf("Queue peek: %s\n", testQueueStr.Peek())

	// Proper Dequeue sequence
	for !testQueueStr.IsEmpty() {
		fmt.Printf("Dequeued: %s\n", testQueueStr.Dequeue())
	}
	fmt.Printf("Queue empty after dequeueing all elements? %v\n", testQueueStr.IsEmpty())

	// Edge Case: Dequeue on empty queue
	if testQueueStr.IsEmpty() {
		fmt.Println("Queue is empty. Cannot dequeue.")
	}

	// Test with custom struct
	fmt.Println("\nTesting Queue with custom struct...")
	type Point struct {
		x, y int
	}
	testQueueStruct := NewQueue[Point]()
	fmt.Printf("Queue empty? %v\n", testQueueStruct.IsEmpty())
	testQueueStruct.Enqueue(Point{1, 2})
	testQueueStruct.Enqueue(Point{3, 4})
	testQueueStruct.Enqueue(Point{5, 6})
	fmt.Printf("Queue size after enqueue: %d\n", testQueueStruct.Size())
	fmt.Printf("Queue peek: %+v\n", testQueueStruct.Peek())

	// Proper Dequeue sequence
	for !testQueueStruct.IsEmpty() {
		fmt.Printf("Dequeued: %+v\n", testQueueStruct.Dequeue())
	}
	fmt.Printf("Queue empty after dequeueing all elements? %v\n", testQueueStruct.IsEmpty())

	// Test edge cases
	fmt.Println("\nTesting edge cases...")
	largeQueue := NewQueue[int]()
	for i := 0; i < 256; i++ { // Test with a large number of elements
		largeQueue.Enqueue(i)
	}
	fmt.Printf("Queue size after enqueuing 256 elements: %d\n", largeQueue.Size())
	for !largeQueue.IsEmpty() {
		largeQueue.Dequeue()
	}
	fmt.Printf("Queue empty after dequeueing all 256 elements? %v\n", largeQueue.IsEmpty())

	// Test IsEmpty on a non-empty queue
	fmt.Println("\nTesting IsEmpty on a non-empty queue...")
	largeQueue.Enqueue(100)
	fmt.Printf("Queue empty? %v\n", largeQueue.IsEmpty())

	fmt.Println("All tests completed successfully.")
}
