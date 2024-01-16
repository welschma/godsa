package ds

// binaryNode is a node in a binary tree.
type binaryNode[K any, V any] struct {
    key K
    value V
    parent *binaryNode[K, V]
    left *binaryNode[K, V]
    right *binaryNode[K, V]
    n int
}

// size returns the number of nodes in the subtree rooted at the node.
func (node *binaryNode[K, V]) size() int {

    if node == nil {
        return 0
    }

    return node.n
}

// deleteMin deletes the minimum node in the subtree rooted at the node.
func (node *binaryNode[K, V]) deleteMin() *binaryNode[K, V] {

    if node.left == nil {
        if node.right != nil {
            node.right.parent = node.parent
        }

        return node.right
    }

    node.left = node.left.deleteMin()
    node.n = node.left.size() + node.right.size() + 1

    return node
}

// deleteMax deletes the maximum node in the subtree rooted at the node.
func (node *binaryNode[K, V]) deleteMax() *binaryNode[K, V] {

    if node.right == nil {
        if node.left != nil {
            node.left.parent = node.parent
        }

        return node.left
    }

    node.right = node.right.deleteMax()

    if node.right != nil {
        node.right.parent = node
    }
    
    node.n = node.left.size() + node.right.size() + 1
    return node
}

// put inserts a new node into the subtree rooted at the node.
func (node *binaryNode[K, V]) put(k K, v V, p *binaryNode[K, V], compareKeyFn func(newKey K, currentKey K) int) *binaryNode[K, V] {

    if node == nil {
        return &binaryNode[K, V]{key: k, value: v, parent: p, left: nil, right: nil, n: 1}
    }

    comp := compareKeyFn(k, node.key)

    if comp < 0 {
        node.left = node.left.put(k, v, node, compareKeyFn)
    } else if comp > 0 {
        node.right = node.right.put(k, v, node, compareKeyFn)
    } else {
        node.value = v
    }

    node.n = node.left.size() + node.right.size() + 1

    return node
}

// get returns the value associated with the given key in the subtree rooted at the node.
func (node *binaryNode[K, V]) get(k K, compareKeyFn func(newKey K, currentKey K) int) (V, bool) {

    if node == nil {
        var v V
        return v, false
    }

    comp := compareKeyFn(k, node.key)

    if comp < 0 {
        return node.left.get(k, compareKeyFn)
    }

    if comp > 0 {
        return node.right.get(k, compareKeyFn)
    }

    return node.value, true
}

// min returns the minimum node in the subtree rooted at the node.
func (node *binaryNode[K, V]) min() *binaryNode[K, V] {

    if node.left == nil {
        return node
    }

    return node.left.min()
}

// delete deletes the node with the given key in the subtree rooted at the node.
func (node *binaryNode[K, V]) delete(k K, compareKeyFn func(newKey K, currentKey K) int) *binaryNode[K, V] {

    if node == nil {
        return nil
    }

    comp := compareKeyFn(k, node.key)

    if comp < 0 {
        node.left = node.left.delete(k, compareKeyFn)
    } else if comp > 0 {
        node.right =  node.right.delete(k, compareKeyFn)
    } else {

        if node.right == nil {

            if node.left != nil {
                node.left.parent = node.parent
            }

            return node.left
        }

        if node.left == nil  {

            if node.right != nil {
                node.right.parent = node.parent
            }
            
            return node.right
        }
        
        t := node
        node = node.right.min()
        node.right = t.right.deleteMin()
        node.left = t.left
        node.parent = t.parent
    }

    node.n = node.left.size() + node.right.size() + 1

    return node

}

// keys returns the keys in the subtree rooted at the node.
func (node *binaryNode[K, V]) keys(keySlice []K) []K {

    if node != nil {
        keySlice = node.left.keys(keySlice)
        keySlice = append(keySlice, node.key)
        keySlice = node.right.keys(keySlice)
    }

    return keySlice
}

// values returns the values in the subtree rooted at the node.
func (node *binaryNode[K, V]) values(valueSlice []V) []V {

    if node != nil {
        valueSlice = node.left.values(valueSlice)
        valueSlice = append(valueSlice, node.value)
        valueSlice = node.right.values(valueSlice)
    }

    return valueSlice
}

// BinarySearchTree is a binary search tree.
type BinarySearchTree[K any, V any] struct {
    root *binaryNode[K, V]
    compareKeyFn func(newKey K, currentKey K) int
}

// Size returns the number of nodes in the tree.
func (bst *BinarySearchTree[K, V]) Size() int {
    return bst.root.size()
}

// IsEmpty returns true if the tree is empty, false otherwise.
func (bst *BinarySearchTree[K, V]) IsEmpty() bool {
    return bst.root == nil
}

// Put inserts a new node into the tree.
func (bst *BinarySearchTree[K, V]) Put(key K, value V) {
    bst.root = bst.root.put(key, value, nil, bst.compareKeyFn)
}

// Get returns the value associated with the given key in the tree.
func (bst *BinarySearchTree[K, V]) Get(key K) (V, bool) {
    return bst.root.get(key, bst.compareKeyFn)
}

// Delete deletes the node with the given key in the tree.
func (bst *BinarySearchTree[K, V]) Delete(key K) {
    bst.root = bst.root.delete(key, bst.compareKeyFn)
}

// Keys returns the keys in the tree.
func (bst *BinarySearchTree[K, V]) Keys() []K {
    return bst.root.keys([]K{})
}

// Values returns the values in the tree.
func (bst *BinarySearchTree[K, V]) Values() []V {
    return bst.root.values([]V{})
}

// DeleteMin deletes the minimum node in the tree.
func (bst *BinarySearchTree[K, V]) DeleteMin() {
    bst.root = bst.root.deleteMin()
}

// DeleteMax deletes the maximum node in the tree.
func (bst *BinarySearchTree[K, V]) DeleteMax() {
    bst.root = bst.root.deleteMax()
}
 
// Clear removes all nodes from the tree.
func (bst *BinarySearchTree[K, V]) Clear() {
    bst.root = nil
}

// NewBinarySearchTree returns a new binary search tree.
func NewBinarySearchTree[K any, V any](compareKeyFn func(K1 K, K2 K) int) BinarySearchTree[K,V] {
    return BinarySearchTree[K, V]{compareKeyFn: compareKeyFn}
}


