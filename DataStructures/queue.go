package DataStructures

// Bare minimum implemenations right now. Could add Doubly ended Queue later but for BFS this is all thats needed
type Queue[T comparable] interface {
	Enqueue(data T) // add elements to front
	Dequeue() T     // Remove an element from the front of the queue
	IsEmpty() bool
	Peek() T // returns front of Q without popping it
	Size() uint8
}
type queue[T comparable] struct {
	elements []T
}

func NewQueue[T comparable]() Queue[T] {
	return newqueueInstance[T]()
}
func newqueueInstance[T comparable]() *queue[T] {
	return &queue[T]{}
}
func (q *queue[T]) Enqueue(data T) {
	q.elements = append(q.elements, data)
}

// Its up to the caller to call IsEmpty before calling Deque or else program will panic
func (q *queue[T]) Dequeue() T {
	data := q.elements[0]
	q.elements = q.elements[1:]
	return data
}
func (q *queue[T]) IsEmpty() bool {
	return len(q.elements) == 0
}
func (q *queue[T]) Peek() T {
	return q.elements[0]
}
func (q *queue[T]) Size() uint8 {
	return uint8(len(q.elements))
}
