package data_structures

import (
	"errors"
	"fmt"
	"strconv"
)

type bt_Node struct {
	value       int64
	left, right *bt_Node
}

type BinaryTree struct {
	head *bt_Node
}

func (bt *bt_Node) add(v int64) error {
	switch {
	case v == bt.value:
		return errors.New(strconv.FormatInt(v, 10) + " already exist in tree")
	case v > bt.value:
		if bt.right == nil {
			bt.right = &bt_Node{value: v}
		} else {
			return bt.right.add(v)
		}
	case v < bt.value:
		if bt.left == nil {
			bt.left = &bt_Node{value: v}
		} else {
			return bt.left.add(v)
		}
	}
	return nil
}
func (bt *BinaryTree) Add(v int64) error {
	if bt.head == nil {
		bt.head = &bt_Node{value: v}
		return nil
	}
	return bt.head.add(v)
}

func (bt bt_Node) printInorder() {
	if bt.left != nil {
		bt.left.printInorder()
	}

	fmt.Println(bt.value)

	if bt.right != nil {
		bt.right.printInorder()
	}
}
func (bt BinaryTree) PrintInorder() error {
	if bt.head == nil {
		return errors.New("tree is empty")
	}
	bt.head.printInorder()
	return nil
}

func (bt bt_Node) printPreorder() {
	fmt.Println(bt.value)

	if bt.left != nil {
		bt.left.printPreorder()
	}

	if bt.right != nil {
		bt.right.printPreorder()
	}
}
func (bt BinaryTree) PrintPreorder() error {
	if bt.head == nil {
		return errors.New("tree is empty")
	}
	bt.head.printPreorder()
	return nil
}

func (bt bt_Node) printPostorder() {
	if bt.left != nil {
		bt.left.printPostorder()
	}

	if bt.right != nil {
		bt.right.printPostorder()
	}

	fmt.Println(bt.value)
}
func (bt BinaryTree) PrintPostorder() error {
	if bt.head == nil {
		return errors.New("tree is empty")
	}
	bt.head.printPostorder()
	return nil
}
