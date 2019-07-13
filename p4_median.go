package leetcode

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	lenX, lenY := len(nums1), len(nums2)
	low, high := 0, lenX

	var maxLeftX, minRigthX, maxLeftY, minRigthY float64
	for low <= high {
		partitionX := (low + high) / 2
		partitionY := (lenX+lenY+1)/2 - partitionX

		if partitionX > 0 {
			maxLeftX = float64(nums1[partitionX-1])
		} else {
			maxLeftX = math.Inf(-1)
		}

		if partitionX == lenX {
			minRigthX = math.Inf(1)
		} else {
			minRigthX = float64(nums1[partitionX])
		}

		if partitionY > 0 {
			maxLeftY = float64(nums2[partitionY-1])
		} else {
			maxLeftY = math.Inf(-1)
		}

		if partitionY == lenY {
			minRigthY = math.Inf(1)
		} else {
			minRigthY = float64(nums2[partitionY])
		}

		if maxLeftX <= minRigthY && maxLeftY <= minRigthX {
			if (lenX+lenY)%2 == 0 {
				return (math.Max(maxLeftX, maxLeftY) + math.Min(minRigthX, minRigthY)) / 2.0
			} else {
				return math.Max(maxLeftX, maxLeftY)
			}
		}
		if maxLeftX > minRigthY {
			high = partitionX - 1
		} else {
			low = partitionX + 1
		}
	}
	return float64(0)
}
