func arithmeticTriplets(nums []int, diff int) (ans int) {
    n := len(nums)
    i, j, k := 0, 1, 2
    for k < n {
        if nums[j] == nums[k] - diff && nums[i] == nums[k] - diff * 2 {
            ans++
            k++
        } else if nums[j] > nums[k] - diff {
            k++
        } else if nums[j] < nums[k] - diff {
            j++
        } else if nums[i] > nums[k] - diff * 2 {
            k++
        } else if nums[i] < nums[k] - diff * 2 {
            i++
        } 
    }

    return
}