package sorter

import (
	"log"
	"reflect"
)

type DataSet []int

func SelectionSort(data DataSet) DataSet {

	for i := 0; i < len(data); i++ {
		min := i
		for j := i + 1; j < len(data); j++ {
			if data[j] < data[min] {
				min = j
			}
		}
		Swap(&data[i], &data[min])
	}
	return data
}

func InsertionSort(data DataSet) DataSet {

	for i := 0; i < len(data) - 1; i++ {
		insert := data[i + 1];
		j := i
		for (j >= 0) && (data[j] > insert) {
			data[j + 1] = data[j]
			j--
		}
		data[j + 1] = insert
	}
	return data

}

func Swap(a interface{}, b interface{}) {

	if _, ok := a.(*int); !ok {
		log.Fatal("Swap: type of a must be *int.")
	}
	if _, ok := b.(*int); !ok {
		log.Fatal("Swap: type of b must be *int.")
	}

	pA := reflect.ValueOf(a).Elem()
	pB := reflect.ValueOf(b).Elem()
	tmp := pA.Interface()

	pA.Set(pB)
	pB.Set(reflect.ValueOf(tmp))
}