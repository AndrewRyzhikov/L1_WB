package main

import "fmt"

func leftBinSearch(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) / 2
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}

	if nums[l] != target {
		return -1
	}

	return l
}

func rightBinSearch(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r + 1) / 2
		if nums[mid] <= target {
			l = mid
		} else {
			r = mid - 1
		}
	}

	if nums[l] != target {
		return -1
	}

	return l
}

func main() {
	a := []int{-1, 0, 3, 5, 9, 12}
	fmt.Println(rightBinSearch(a, 9))
	fmt.Println(leftBinSearch(a, 9))
}
