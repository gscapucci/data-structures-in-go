package data_structures

import (
	"errors"
	"fmt"
	"strconv"
)

type Stack struct {
	values []int64
}

func (stack *Stack) Push(val int64) {
	stack.values = append(stack.values, val)
}

func (stack *Stack) Pop() error {
	if len(stack.values) <= 0 {
		return errors.New("stack is empty")
	}
	stack.values = stack.values[:len(stack.values)-1]
	return nil
}

func (stack Stack) PrintStack() error {
	if len(stack.values) <= 0 {
		return errors.New("list is empty")
	}

	for i := 0; i < len(stack.values); i++ {
		fmt.Print(stack.values[i], " ")
	}
	return nil
}

func (stack *Stack) Get() (int64, error) {
	if len(stack.values) <= 0 {
		return 0, errors.New("list is empty")
	}
	num := stack.values[len(stack.values)-1]
	stack.Pop()
	return num, nil
}

func (stack Stack) Peek() (int64, error) {
	if len(stack.values) == 0 {
		return 0, errors.New("stack is empty")
	}
	return stack.values[len(stack.values)-1], nil
}

func (stack Stack) Search(val int64) (int64, error) {
	if len(stack.values) == 0 {
		return 0, errors.New("list is empty")
	}
	for i := 0; i < len(stack.values); i++ {
		if stack.values[i] == val {
			return stack.values[i], nil
		}
	}
	return 0, errors.New(strconv.FormatInt(val, 10) + " was not found")
}

func (stack Stack) IsEmpty() bool {
	return len(stack.values) == 0
}
