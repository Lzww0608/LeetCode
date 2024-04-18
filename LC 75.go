func sortColors(nums []int)  {
    n := len(nums)
    i, j, k := 0, 0, n - 1

    for i <= k {
        if nums[i] == 2 {
            nums[i], nums[k] = nums[k], nums[i]
            k--
        } else if nums[i] == 0 {
            nums[j], nums[i] = nums[i], nums[j]
            i++
            j++
        } else {
            i++
        }
        //fmt.Println(i, j, k, nums)
    }

    return
}