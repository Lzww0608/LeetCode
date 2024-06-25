func moveZeroes(nums []int)  {
    n := len(nums)
    l, r := 0, 0
    for l, r = 0, 0; r < n; r++ {
        if nums[r] != 0 {
            nums[l] = nums[r]
            l++
        }
    }

    for l < n {
        nums[l] = 0
        l++
    }

}
