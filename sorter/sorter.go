package sorter

import (
	"log"
	"reflect"
)

type DataSet []int

const c = 2

func SelectionSort(data DataSet) DataSet {

	for i := 0; i < len(data); i++ {
		min := i
		for j := i + 1; j < len(data); j++ {
			if data[j] < data[min] {
				min = j
			}
		}
		swap(&data[i], &data[min])
	}
	return data
}

func InsertionSort(data DataSet) DataSet {

	for i := 0; i < len(data) - 1; i++ {
		insert := data[i + 1]
		j := i
		for (j >= 0) && (data[j] > insert) {
			data[j + 1] = data[j]
			j--
		}
		data[j + 1] = insert
	}
	return data

}

func ShellSort(data DataSet) DataSet {

	m := len(data) / 2
	for m > 0 {
		for i := 0; i < len(data) - m; i++ {
			insert := data[m + i]
			j := i
			for (j >= 0) && (data[j] > insert) {
				data[m + j] = data[j]
				j = j - m
			}
			data[j + m] = insert
		}
		m = m / 2
	}
	return nil

}

func QuickSort(data *DataSet) {

	if len(*data) > 1 {
		i, j := 0, len(*data) - 1
		x := midValue((*data)[0], (*data)[(len(*data) - 1) / 2], (*data)[len(*data) - 1])
		for i <= j {
			for (*data)[i] <= x {
				i++
			}

			for (*data)[j] > x {
				j--
			}

			if i <= j {
				swap(&(*data)[i], &(*data)[j])
				i++
				j--
			}
		}
		left, right := (*data)[:j], (*data)[i:]
		QuickSort(&left)
		QuickSort(&right)
	}

}

func midValue(x, y, z int) int {

	l, m, h := x, y, z
	if h < l {
		swap(&h, &l)
	}
	if l > m {
		swap(&l, &m)
	} else if h < m {
		swap(&h, &m)
	}
	return m
}

func swap(a interface{}, b interface{}) {

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