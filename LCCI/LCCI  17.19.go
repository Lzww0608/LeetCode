func missingTwo(nums []int) (ans []int) {
    nums = append(nums, -1, -1)
    for i := range nums {
        for i != nums[i] - 1 && nums[i] != -1 {
            nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
        }
    }

    for i, x := range nums {
        if x == -1 {
            ans = append(ans, i + 1)
        }
    }

    return 
}