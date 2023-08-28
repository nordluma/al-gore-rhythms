package main

type SinglyLinkedNode[T any] struct {
	Value T
	Next  *SinglyLinkedNode[T]
}

func newSinglyLinkedNode[T any](value T) *SinglyLinkedNode[T] {
	return &SinglyLinkedNode[T]{Value: value, Next: nil}
}

type DoublyLinkedNode[T any] struct {
	Value T
	Next  *DoublyLinkedNode[T]
	Prev  *DoublyLinkedNode[T]
}

func newDoublyLinkedNode[T any](value T) *DoublyLinkedNode[T] {
	return &DoublyLinkedNode[T]{Value: value, Next: nil, Prev: nil}
}
