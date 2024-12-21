package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// O(n+m)
	n, m, a, b := len(nums1), len(nums2), 0, 0
	last := make([]int, n+m)

	for i := 0; i < n+m; i++ {
		if a >= n {
			last[i] = nums2[b]
			b++
			continue
		} else if b >= m {
			last[i] = nums1[a]
			a++
			continue
		}

		if nums1[a] < nums2[b] {
			last[i] = nums1[a]
			a++
		} else if nums1[a] == nums2[b] {
			last[i] = nums1[a]
			a++
		} else {
			last[i] = nums2[b]
			b++
		}
	}

	if (n+m)%2 == 0 {
		return float64(last[(n+m)/2]+last[(n+m)/2-1]) / 2.0
	} else {
		return float64(last[(n+m)/2])
	}
}
