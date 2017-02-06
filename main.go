package main

import (
	"fmt"
	"sorted/sorter"
)

func main() {

	arr := make(sorter.DataSet, 10)
	for k, _  := range arr {
		arr[k] = 10 - k
	}

	fmt.Printf("\nSelectionSort\nUnsorted: %v\n", arr)
	sorter.SelectionSort(arr)
	fmt.Printf("Sorted:   %v\n", arr)

	for k, _  := range arr {
		arr[k] = 10 - k
	}

	fmt.Printf("\nInsertionSort\nUnsorted: %v\n", arr)
	sorter.InsertionSort(arr)
	fmt.Printf("Sorted:   %v\n", arr)

	for k, _  := range arr {
		arr[k] = 10 - k
	}

	fmt.Printf("\nShellSort\nUnsorted: %v\n", arr)
	sorter.ShellSort(arr)
	fmt.Printf("Sorted:   %v\n", arr)

	for k, _  := range arr {
		arr[k] = 10 - k
	}

	fmt.Printf("\nQuickSort\nUnsorted: %v\n", arr)
	sorter.QuickSort(&arr)
	fmt.Printf("Sorted:   %v\n", arr)

}
