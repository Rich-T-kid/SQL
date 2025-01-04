package DataStructures

import "fmt"

// Exact same as constraints.Ordered but Doesnt require me to update the go Version
type Ordered interface {
	~int | ~float64 | ~string
}

type node[T Ordered] struct {
	value T
	next  *node[T]
	prev  *node[T]
}

type LinkedList[T Ordered] struct {
	head *node[T]
	tail *node[T]
}

type Chainer[T Ordered] interface {
	// Add an element while maintaining sorted order.
	//Insert(data T) // O(n)
	Append(data T)
	// Remove the first (head) or last (tail) element.
	PopHead() *node[T]
	PopTail() *node[T]

	// Find an element (useful for search operations in B+ Trees).
	Search(data T) *node[T] // log n Time due to sorted property

	// Merge two sorted linked lists into one sorted list (splitting or merging B+ Tree nodes).
	Merge(other *LinkedList[T]) *node[T]

	Display()
	// Split the linked list into two parts at the given position (splitting B+ Tree nodes).
	Split(position int) (Chainer[T], Chainer[T])

	// Get the size of the linked list.
	Size() uint32

	Sort()

	// Check if the linked list is empty.
	Empty() bool
}

func newChainer[T Ordered]() Chainer[T] {
	return newLinkedList[T]()
}

// ---------------------------- //
//       Constructors           //
// ---------------------------- //

func newLinkedList[T Ordered]() *LinkedList[T] {
	return &LinkedList[T]{}
}
func (l *LinkedList[T]) Split(position int) (Chainer[T], Chainer[T]) {
	return nil, nil
}

func newNode[T Ordered](data T) *node[T] {
	return &node[T]{
		value: data,
	}
}

// These should not throw Errors. in case of client errors such as attempting to pop from ann empty linked list nothing should happen
func (l *LinkedList[T]) Addfront(data T) {
	// Case where linked list is empty
	newNode := newNode(data)
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}
	// Case wherre linked List is not empty
	l.head.prev = newNode
	newNode.next = l.head
	l.head = newNode
}

// Add to end of Linked list
func (l *LinkedList[T]) Insert(data T) {

}
func (l *LinkedList[T]) Append(data T) {
	// case where linked list is empty
	newNode := newNode(data)
	if l.head == nil {
		// also no tail so init the tail and the head and set the values
		l.head = newNode
		l.tail = newNode
		return
	}
	l.tail.next = newNode
	newNode.prev = l.tail
	l.tail = newNode

	// case where linked list is not empty
}

// Pop head and returns its value. Up to caller to decide if they wish to use the value
func (l *LinkedList[T]) PopHead() *node[T] {
	// case where linked Lisst is empty
	if l.head == nil {
		return nil
	}
	// case where theres only one element
	if l.head.next == nil {
		node := l.head
		l.head = nil
		l.tail = nil
		return node
	}
	// case where the head has a succsessor
	oldhead := l.head
	newhead := l.head.next
	newhead.prev = nil // disconnect from prevous head
	oldhead.next = nil
	l.head = newhead
	return oldhead
}
func (l *LinkedList[T]) PopTail() *node[T] {
	// Case where list is empty
	if l.tail == nil {
		return nil
	}
	// case where its only one element
	if l.tail.prev == nil {
		node := l.tail
		l.head = nil
		l.tail = nil
		return node
	}
	// case where the node has a successor
	oldtail := l.tail
	newtail := l.tail.prev
	oldtail.prev = nil // disconnect from previous tail
	newtail.next = nil // make new tail
	l.tail = newtail
	return oldtail
}

// Prints 10 nodes per line before ending of with a messsage of End of Linked List
func (l *LinkedList[T]) Display() {
	current := l.head
	var count int
	var line string
	// 10 nodes per line
	for current != nil {
		str := fmt.Sprintf("node: %v, ", current.value)
		line += str
		count++
		if count > 0 && count%10 == 0 {
			fmt.Println(line)
			line = " "
			count = 0
		}
		current = current.next
	}
	fmt.Println("End Of Linked List")
}

// Searches for data of type T, if node with data is found it returns the node that holds this data
func (l *LinkedList[T]) Search(data T) *node[T] {
	current := l.head
	for current != nil {
		if current.value == data {
			return current
		} else {
			current = current.next
		}
	}
	return nil
}
func (l *LinkedList[T]) Clear() {
	l.head = nil
	l.tail = nil
}
func (l *LinkedList[T]) Size() uint32 {
	var size uint32
	curr := l.head
	for curr != nil {
		size++
		curr = curr.next
	}
	return size
}
func (l *LinkedList[T]) Sort() {
	newhead := l.MergeSort()
	l.head = newhead

}

func (l *LinkedList[T]) MergeSort() *node[T] {
	// Base case
	fmt.Printf("[DEBUG] MergeSort called. Size=%d\n", l.Size())
	if l.head == nil || l.head.next == nil {
		return l.head
	}

	mid := l.middle()

	rightList := &LinkedList[T]{head: mid.next}
	if rightList.head != nil {
		rightList.head.prev = nil
	}
	mid.next = nil // Break into two halves

	leftList := &LinkedList[T]{head: l.head}
	leftSorted := leftList.MergeSort()
	rightSorted := rightList.MergeSort()

	mergedList := &LinkedList[T]{}
	mergedList.head = (&LinkedList[T]{head: leftSorted}).Merge(
		&LinkedList[T]{head: rightSorted},
	)
	return mergedList.head
}

// Merge nodes in sorted order
// Assumes that the inputed Linked List are sorted
func (l *LinkedList[T]) Merge(other *LinkedList[T]) *node[T] {
	dummy := &node[T]{}
	current := dummy

	p1 := l.head
	p2 := other.head

	for p1 != nil && p2 != nil {
		if p1.value <= p2.value {
			current.next = p1
			p1.prev = current
			p1 = p1.next
		} else {
			current.next = p2
			p2.prev = current
			p2 = p2.next
		}
		current = current.next
	}

	// Append the remaining nodes from either list
	if p1 != nil {
		current.next = p1
		p1.prev = current // <-- fix leftover p1 node
	}
	if p2 != nil {
		current.next = p2
		p2.prev = current // <-- fix leftover p2 node
	}

	newHead := dummy.next
	if newHead != nil {
		newHead.prev = nil
	}

	return newHead
}

func (l *LinkedList[T]) middle() *node[T] {
	if l.head == nil {
		return nil
	}
	fast := l.head.next
	slow := l.head
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
	}
	return slow
}
func (l *LinkedList[T]) hasCycle() bool {
	slow, fast := l.head, l.head
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
		if fast == slow {
			return true
		}
	}
	return false
}

func (l *LinkedList[T]) ToSlice() []T {
	var slice []T
	curr := l.head
	for curr != nil {
		slice = append(slice, curr.value)
		curr = curr.next
	}
	return slice
}

// Append inputed linked List to end of current Linked List
// ------------------------------------- //
//           Concatenation              //
// ------------------------------------- //

// CHANGED: Update 'l.tail' after concatenation so it's correct
func (l *LinkedList[T]) Concate(linkedList *LinkedList[T]) []T {
	if l.head == nil {
		// Inherit both head & tail from linkedList
		l.head = linkedList.head
		l.tail = linkedList.tail
		return l.ToSlice()
	}
	if linkedList.head == nil {
		return l.ToSlice()
	}

	// Find the tail of the current list
	current := l.tail // CHANGED: we already track tail, no need to traverse
	current.next = linkedList.head
	linkedList.head.prev = current

	// ADDED: Now update 'l.tail' to be the tail of 'linkedList'
	if linkedList.tail != nil {
		l.tail = linkedList.tail
	} else {
		// If linkedList had only one node, now that node is also our tail
		l.tail = linkedList.head
	}
	return l.ToSlice()
}

func (l *LinkedList[T]) Empty() bool {
	return l.Size() == 0
}
