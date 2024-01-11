package main

import "fmt"

// array
type Array struct {
	data []interface{}
	size int
  }
  
  func NewArray(size int) *Array {
	return &Array{
	  data: make([]interface{}, size),
	  size: size,
	}
  }
  
  func (a *Array) Get(index int) interface{} {
	if index < 0 || index >= a.size {
	  panic("Index out of bounds")
	}
	return a.data[index]
  }
  
  func (a *Array) Set(index int, value interface{}) {
	if index < 0 || index >= a.size {
	  panic("Index out of bounds")
	}
	a.data[index] = value
  }

//   linkedlist
type Node struct {
	data interface{}
	next *Node
  }
  
  type LinkedList struct {
	head *Node
	tail *Node
	size int
  }
  
  func NewLinkedList() *LinkedList {
	return &LinkedList{
	  head: nil,
	  tail: nil,
	  size: 0,
	}
  }
  
  func (ll *LinkedList) AddFirst(data interface{}) {
	newNode := &Node{data: data}
	if ll.head == nil {
	  ll.head = newNode
	  ll.tail = newNode
	} else {
	  newNode.next = ll.head
	  ll.head = newNode
	}
	ll.size++
  }
  
  func (ll *LinkedList) AddLast(data interface{}) {
	newNode := &Node{data: data}
	if ll.tail == nil {
	  ll.head = newNode
	  ll.tail = newNode
	} else {
	  ll.tail.next = newNode
	  ll.tail = newNode
	}
	ll.size++
  }
  
  func (ll *LinkedList) RemoveFirst() interface{} {
	if ll.head == nil {
	  panic("Empty list")
	}
	data := ll.head.data
	ll.head = ll.head.next
	if ll.head == nil {
	  ll.tail = nil
	}
	ll.size--
	return data
  }

//   Stack
type Stack struct {
	top *Node
	size int
  }
  
  func NewStack() *Stack {
	return &Stack{
	  top: nil,
	  size: 0,
	}
  }
  
  func (s *Stack) Push(data interface{}) {
	newNode := &Node{data: data}
	newNode.next = s.top
	s.top = newNode
	s.size++
  }
  
  func (s *Stack) Pop() interface{} {
	if s.top == nil {
	  panic("Empty stack")
	}
	data := s.top.data
	s.top = s.top.next
	s.size--
	return data
  }

//   Queue
type Queue struct {
	front *Node
	rear *Node
	size int
  }
  
  func NewQueue() *Queue {
	return &Queue{
	  front: nil,
	  rear: nil,
	  size: 0,
	}
  }
  
  func (q *Queue) Enqueue(data interface{}) {
	newNode := &Node{data: data}
	if q.rear == nil {
	  q.front = newNode
	  q.rear = newNode
	} else {
	  q.rear.next = newNode
	  q.rear = newNode
	}
	q.size++
  }
  
  func (q *Queue) Dequeue() interface{} {
	if q.front == nil {
	  panic("Empty queue")
	}
	data := q.front.data
	q.front = q.front.next
	if q.front == nil {
	  q.rear = nil
	}
	q.size--
	return data
  }

  
//   BinarySearch Tree
type TreeNode struct {
	data interface{}
	left *TreeNode
	right *TreeNode
  }
  
  type BinarySearchTree struct {
	root *TreeNode
  }
  
  func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{
	  root: nil,
	}
  }
  
  func (bst *BinarySearchTree) Insert(data interface{}) {
	bst.root = bst.insert(bst.root, data)
  }
  
  func (bst *BinarySearchTree) Search(data interface{}) bool {
	return bst.search(bst.root, data)
  }
  
  func (bst *BinarySearchTree) Delete(data interface{}) {
	bst.root = bst.delete(bst.root, data)
  }
  
  func (bst *BinarySearchTree) insert(node *TreeNode, data interface{}) *TreeNode {
	if node == nil {
	  return &TreeNode{data: data}
	}
	if data < node.data {
	  node.left = bst.insert(node.left, data)
	} else {
	  node.right = bst.insert(node.right, data)
	}
	return node
  }
  
  func (bst *BinarySearchTree) search(node *TreeNode, data interface{}) bool {
	if node == nil {
	  return false
	}
	if data == node.data {
	  return true
	} else if data < node.data {
	  return bst.search(node.left, data)
	} else {
	  return bst.search(node.right, data)
	}
  }
  
  func (bst *BinarySearchTree) delete(node *TreeNode, data interface{}) *TreeNode {
	if node == nil {
	  return nil
	}
	if data == node.data {
	  if node.left == nil && node.right == nil {
		return nil
	  } else if node.left == nil {
		return node.right
	  } else if node.right == nil {
		return node.left
	  } else {
		min := bst.min(node.right)
		node.data = min.data
		node.right = bst.delete(node.right, min.data)
		return node
	  }
	} else if data < node.data {
	  node.left = bst.delete(node.left, data)
	} else {
	  node.right = bst.delete(node.right, data)
	}
	return node
  }
  
  func (bst *BinarySearchTree) min(node *TreeNode) *TreeNode {
	if node == nil || node.left == nil {
	  return node
	}
	return bst.min(node.left)
  }
  

//   Graph
type Graph struct {
	vertices int
	edges map[int][]int
  }
  
  func NewGraph(vertices int) *Graph {
	graph := &Graph{
	  vertices: vertices,
	  edges: make(map[int][]int),
	}
	for i := 0; i < vertices; i++ {
	  graph.edges[i] = []int{}
	}
	return graph
  }
  
  func (g *Graph) AddEdge(source, destination int) {
	g.edges[source] = append(g.edges[source], destination)
  }
  
  func (g *Graph) PrintAdjacencyList() {
	for vertex, edges := range g.edges {
	  fmt.Printf("Vertex %d: ", vertex)
	  for _, edge := range edges {
		fmt.Printf("%d ", edge)
	  }
	  fmt.Println()
	}
  }
  