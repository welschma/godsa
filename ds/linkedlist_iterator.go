package ds

type LinkedListIterator[T comparable] struct {
    header *listNode[T]
    currentNode *listNode[T]
}


func (it *LinkedListIterator[T]) GetNext() T {
    var t T

    if it.HasNext() {
        t = it.currentNode.val
        it.currentNode = it.currentNode.next
        return t 
    }

    return t
}

func (it *LinkedListIterator[T]) HasNext() bool {
    return it.currentNode != it.header
}
