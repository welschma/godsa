package list

type SinglyLinkedNode[T comparable] struct {
	item T
	next *SinglyLinkedNode[T]
}

type SinglyLinkedList[T comparable] struct {
	head *SinglyLinkedNode[T]
	size int
}

func NewSinglyLinkedList[T comparable]() SinglyLinkedList[T] {
	return SinglyLinkedList[T]{}
}

func (sll *SinglyLinkedList[T]) IsEmpty() bool {
	return (sll.head == nil)
}

func (sll *SinglyLinkedList[T]) Size() int {
	return sll.size
}

func (sll *SinglyLinkedList[T]) initEmpty(newNode *SinglyLinkedNode[T]) {
	sll.head = newNode
}

func (sll *SinglyLinkedList[T]) InsertAtTail(newItem T) {
	newNode := SinglyLinkedNode[T]{item: newItem}
	sll.size += 1

	if sll.IsEmpty() {
		sll.initEmpty(&newNode)
		return
	}

	lastNode := sll.head

	for lastNode.next != nil {
		lastNode = lastNode.next
	}

	lastNode.next = &newNode
}

func (sll *SinglyLinkedList[T]) InsertAtHead(newItem T) {
	newNode := SinglyLinkedNode[T]{item: newItem}
	sll.size += 1

	if sll.IsEmpty() {
		sll.initEmpty(&newNode)
		return
	}

	newNode.next = sll.head
	sll.head = &newNode
}

func (sll *SinglyLinkedList[T]) DeleteFromHead() (T, error) {

	if sll.IsEmpty() {
		var zeroValue T
		return zeroValue, ErrEmptyList
	}

	nodeToDelete := sll.head
	sll.head = nodeToDelete.next
	sll.size -= 1
	return nodeToDelete.item, nil
}

func (sll *SinglyLinkedList[T]) DeleteFromTail() (T, error) {

	if sll.IsEmpty() {
		var zeroValue T
		return zeroValue, ErrEmptyList
	}

	if sll.head.next == nil {
		nodeToDelete := sll.head
		sll.head = nil
		sll.size -= 1
		return nodeToDelete.item, nil
	}

	secondToLastNode := sll.head
	for secondToLastNode.next.next != nil {
		secondToLastNode = secondToLastNode.next
	}

	nodeToDelete := secondToLastNode.next
	secondToLastNode.next = nil
	sll.size -= 1
	return nodeToDelete.item, nil
}


func (sll *SinglyLinkedList[T]) Search(data T) bool {
	currentNode := sll.head

	if sll.IsEmpty() {
		return false
	}
	
	if currentNode.item == data {
			return true
		}

	for currentNode.next != nil {
		currentNode = currentNode.next
		if currentNode.item == data {
			return true
		}
	}

	return false

}

func (sll *SinglyLinkedList[T]) Delete(data T) bool {

	if sll.IsEmpty() {
		return false
	}
	
	if sll.head.item == data {
			sll.head = sll.head.next
			return true
		}

	currentNode := sll.head

	tmp := new(SinglyLinkedNode[T])
	
	for currentNode.next != nil {

		if currentNode.next.item == data {
			tmp = currentNode.next.next
			currentNode.next = tmp
			return true
		}

		currentNode = currentNode.next
	}

	return false

}

func (sll *SinglyLinkedList[T]) ToArray() []T {
	currentNode := sll.head

	if sll.IsEmpty() {
		return []T{}
	}
	
	values := make([]T, 0)
	for currentNode != nil {
		values = append(values, currentNode.item)
		currentNode = currentNode.next
	}

	return values

}