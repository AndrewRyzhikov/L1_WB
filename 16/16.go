package main

import "fmt"

func Partition(v []int, l int, r int) int {
	pivot := (v[l] + v[r]) / 2

	for l <= r {
		for v[l] < pivot {
			l++
		}
		for v[r] > pivot {
			r--
		}

		if l >= r {
			break
		}
		v[l], v[r] = v[r], v[l]
		l++
		r--
	}

	return r
}

func QuickSort(v []int, l int, r int) {
	if l == r {
		return
	}
	partition := Partition(v, l, r)
	QuickSort(v, l, partition)
	QuickSort(v, partition+1, r)
}

func main() {
	v := []int{1, 3, 3, 1, -1, 4, 10, 414, 0, 51}
	QuickSort(v, 0, len(v)-1)
	fmt.Println(v)
}
