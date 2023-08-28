package main

type SinglyLinkedNode[T any] struct {
	Value T
	Next  *SinglyLinkedNode[T]
}

func newNode[T any](value T) *SinglyLinkedNode[T] {
	return &SinglyLinkedNode[T]{Value: value, Next: nil}
}
