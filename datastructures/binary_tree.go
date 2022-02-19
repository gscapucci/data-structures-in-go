package datastructures

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

func (bt *bt_Node) add(val int64) error {
	switch {
	case val == bt.value:
		return errors.New(strconv.FormatInt(val, 10) + " already exist in tree")
	case val > bt.value:
		if bt.right == nil {
			bt.right = &bt_Node{value: val}
		} else {
			return bt.right.add(val)
		}
	case val < bt.value:
		if bt.left == nil {
			bt.left = &bt_Node{value: val}
		} else {
			return bt.left.add(val)
		}
	}
	return nil
}

func (bt *BinaryTree) Add(val int64) error {
	if bt.head == nil {
		bt.head = &bt_Node{value: val}
		return nil
	}
	return bt.head.add(val)
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

func (bt *bt_Node) remove(val int64) error {
	if bt.value == val {
		if bt.left == nil && bt.right == nil {
			bt = nil
		} else {
			v, err := bt.findSub()
			if err != nil {
				return err
			}
			bt.remove(v)
			bt.value = v
			return nil
		}
	} else if val > bt.value {
		err := bt.right.remove(val)
		return err
	} else {
		err := bt.left.remove(val)
		return err
	}
	return errors.New("idk")
}

func (bt bt_Node) getMin() int64 {
	for bt.left != nil {
		bt = *bt.left
	}
	return bt.value
}

func (bt BinaryTree) GetMin() (int64, error) {
	if bt.head == nil {
		return 0, errors.New("tree is empty")
	}
	return bt.head.getMin(), nil
}

func (bt bt_Node) getMax() int64 {
	for bt.right != nil {
		bt = *bt.right
	}
	return bt.value
}

func (bt BinaryTree) GetMax() (int64, error) {
	if bt.head == nil {
		return 0, errors.New("tree is empty")
	}
	return bt.head.getMax(), nil
}

func (bt bt_Node) findSub() (int64, error) {
	if bt.left == nil && bt.right == nil {
		return 0, errors.New("can not find sub node, this node is a leaf")
	}
	var maxLeft, minRight *int64
	if bt.right != nil {
		minRight = new(int64)
		*minRight = bt.right.getMin()
	}
	if bt.left != nil {
		maxLeft = new(int64)
		*maxLeft = bt.left.getMax()
	}
	if maxLeft != nil && minRight != nil {
		var diffR, diffL *int64
		diffR = new(int64)
		diffL = new(int64)
		*diffR = *minRight - bt.value
		*diffL = bt.value - *maxLeft
		if *diffL < *diffR {
			return *maxLeft, nil
		} else {
			return *minRight, nil
		}
	} else if maxLeft == nil {
		return *minRight, nil
	} else {
		return *maxLeft, nil
	}
}

func (bt *BinaryTree) Remove(val int64) error {
	if bt.head == nil {
		return errors.New("tree is empty")
	}
	return bt.head.remove(val)
}
