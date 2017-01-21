package main

import (
	"fmt"
	"sorted/sorter"
)

func main() {

	arr := make([]int, 10)
	for k, _  := range arr {
		arr[k] = 10 - k
	}

	fmt.Printf("%v\n", arr)
	sorter.SelectionSort(arr)
	fmt.Printf("%v\n", arr)

	for k, _  := range arr {
		arr[k] = 10 - k
	}

	fmt.Printf("%v\n", arr)
	sorter.InsertionSort(arr)
	fmt.Printf("%v\n", arr)

}
