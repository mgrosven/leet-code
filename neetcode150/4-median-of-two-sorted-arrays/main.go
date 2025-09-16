package main

import (
	"fmt"
	"math"
)

func main() {
	nums1 := []int{}
	nums2 := []int{1}
	o := findMedianSortedArrays(nums1, nums2)
	fmt.Println(o)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m, n := len(nums1), len(nums2)

	low, high := 0, m
	half := (m + n + 1) / 2

	for low <= high {
		partX := (low + high) / 2
		partY := half - partX

		maxX := math.MinInt64
		if partX > 0 {
			maxX = nums1[partX-1]
		}

		nextX := math.MaxInt64
		if partX < m {
			nextX = nums1[partX]
		}

		maxY := math.MinInt64
		if partY > 0 {
			maxY = nums2[partY-1]
		}
		nextY := math.MaxInt64
		if partY < n {
			nextY = nums2[partY]
		}

		if maxX <= nextY && maxY <= nextX {
			if (m+n)%2 == 0 {
				return (float64(max(maxX, maxY)) + float64(min(nextX, nextY))) / 2
			}
			return float64(max(maxX, maxY))
		} else if maxX > nextY {
			high = partX - 1
		} else {
			low = partX + 1
		}
	}
	return 0.0
}
