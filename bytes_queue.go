package BytesQueue

type BytesQueue struct {
	full			 bool
	capacity         int
	maxCapacity      int
	count            int
	rightMarginIndex int
	head             int
	tail             int
	verbose          bool
	array            []byte
	headerBuffer     []byte
}

type queueError struct {
	message string
}

var (
	errEmptyQueue       = &queueError{message: "Empty Queue"}
	errInvalidIndex     = &queueError{message: "Index must be greater than zero. Invalid Index."}
	errIndexOutOfBounds = &queueError{message: "Index out of range."}
)

//======================================== Exported Methods =====================================//

func (q *BytesQueue) Push(data []byte) (index int, err error) {
	// Obtain the number of bytes needed to write the length of data.
	// Check if can't add after tail and depending upon that take appropriate route.
	// In the end return the index at which the item is added.
	// Also send the error.
	return
}

func (q * BytesQueue) Peek() (data []byte, err error) {
	// Peek the item that may be popped.
	return
}

func (q *BytesQueue) Pop() (data []byte, err error) {
	// Pop the item from the queue.
	return
}

func (q *BytesQueue) Capacity() int {
	return q.capacity
}

func (q *BytesQueue) Get(index int) (data []byte, err error) {
	// Return the item at particular index.
	return
}

func (q *BytesQueue) Len() int {
	return q.count
}

func (q *BytesQueue) CheckGet(index int) (err error) {
	// Check if an entry can be read from index
	return
}

func (e *queueError) Error() string {
	return e.message
}

//======================================== Helper Methods =======================================//




