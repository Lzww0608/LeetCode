func sortArray(nums []int) []int {
    n := len(nums)
    
    var quickSort func(int, int) 
    quickSort = func(l, r int) {
        if l >= r {
            return 
        }
        pivot := nums[l + rand.Intn(r - l)] 
        //pivot := nums[l]
        i, j, k := l, l, r + 1
        for i < k {
            if nums[i] < pivot {
                nums[i], nums[j] = nums[j], nums[i]
                i++
                j++
            } else if nums[i] > pivot {
                k--
                nums[i], nums[k] = nums[k], nums[i] 
            } else {
                i++
            }
        }

        quickSort(l, j - 1)
        quickSort(k, r)
    }

    quickSort(0, n - 1)
    return nums
}
