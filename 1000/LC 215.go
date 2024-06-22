func findKthLargest(nums []int, k int) int {
    n := len(nums)
    rand.Seed(time.Now().UnixNano())

    var quickSort func(int, int) int
    quickSort = func(l, r int) int {
        if l >= r {
            return nums[l]
        }
        pivot := nums[l + rand.Intn(r - l + 1)]
        i, j, m := l, l, r + 1
        for i < m {
            if nums[i] > pivot {
                m--
                nums[i], nums[m] = nums[m], nums[i]
            } else if nums[i] < pivot {
                nums[i], nums[j] = nums[j], nums[i]
                i++
                j++
            } else {
                i++
            }
        }

        if n - j < k {
			return quickSort(l, j-1)
		} else if n - m >= k {
			return quickSort(m, r)
		}
		return nums[j]
    }

    return quickSort(0, n - 1)
}
