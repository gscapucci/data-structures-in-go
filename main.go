package main

import (
	"fmt"

	ds "github.com/gscapucci/data-structures-in-go/data_structures"
)

func main() {
	var list ds.LinkedList
	for i := int64(0); i < 10; i++ {
		list.Push_back(i)
	}

	_, err := list.Get(1000)
	if err != nil {
		fmt.Println(err.Error())
	}
}
