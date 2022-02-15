package data_structures

import (
	"errors"
	"fmt"
	"strconv"
)

type ll_Node struct {
	value int64
	next  *ll_Node
}
type LinkedList struct {
	head *ll_Node
	size uint64
}

func (list *LinkedList) Push_back(v int64) {
	if list.head == nil {
		list.head = &ll_Node{value: v}
		list.size++
		return
	}

	var node *ll_Node = list.head
	for node.next != nil {
		node = node.next
	}
	node.next = &ll_Node{value: v}
	list.size++
}

func (list *LinkedList) Push_front(v int64) {
	if list.head == nil {
		list.head = &ll_Node{value: v}
		list.size++
		return
	}

	var node *ll_Node = &ll_Node{value: v}
	node.next = list.head
	list.head = node.next
	list.size++
}

func (list *LinkedList) Pop_back() error {
	if list.head == nil {
		return errors.New("list in empty")
	}
	if list.head.next == nil {
		list.head = nil
		list.size--
		return nil
	}
	var aux *ll_Node = list.head
	for aux.next.next != nil {
		aux = aux.next
	}
	aux.next = nil
	list.size--
	return nil
}

func (list *LinkedList) Pop_front() error {
	if list.head == nil {
		return errors.New("list is empty")
	}
	list.head = list.head.next
	list.size--
	return nil
}

func (list LinkedList) Get(index uint64) (int64, error) {
	if list.head == nil {
		return 0, errors.New("list is empty")
	}
	if index >= list.size {
		return 0, errors.New("index " + strconv.FormatUint(index, 10) + " out of range, max index: " + strconv.FormatUint(list.size-1, 10))
	}
	var aux *ll_Node = list.head
	for i := uint64(0); i < index; i++ {
		aux = aux.next
	}
	return aux.value, nil
}

func (list LinkedList) Print() error {
	if list.head == nil {
		return errors.New("list in empty")
	}

	var ll *ll_Node = list.head

	for ll != nil {
		fmt.Print(ll.value, " ")
		ll = ll.next
	}
	fmt.Println()
	return nil
}

func (list LinkedList) Size() uint64 {
	return list.size
}

func (list *LinkedList) Remove(index uint64) error {
	if list.head == nil {
		return errors.New("list is empty")
	}
	if index >= list.size {
		return errors.New("index " + strconv.FormatUint(index, 10) + " out of range, max index: " + strconv.FormatUint(list.size-1, 10))
	}
	if index == 0 {
		return list.Pop_front()
	}
	if index == list.size-1 {
		return list.Pop_back()
	}
	var aux *ll_Node = list.head
	for i := uint64(0); i < index-1; i++ {
		aux = aux.next
	}
	aux.next = aux.next.next
	return nil
}

func (list *LinkedList) Reverse() error {
	if list.head == nil {
		return errors.New("list is empty")
	}
	if list.size == 1 {
		return nil
	}
	var prev, curr, next *ll_Node
	curr = list.head
	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	list.head = prev
	return nil
}

func (list *LinkedList) Clear() error {
	if list.head == nil {
		return errors.New("list in empty")
	}
	list.head = nil
	return nil
}
