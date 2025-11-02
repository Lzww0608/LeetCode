func getSneakyNumbers(nums []int) (ans []int) {
    n := len(nums)
    for i := range n {
        if i == nums[i] {
            continue
        }

        if nums[nums[i]] == nums[i] {
            ans = append(ans, nums[i])
            nums[i] = -1
            continue
        }

        for nums[i] >= 0 && nums[i] != i {
            if nums[i] == nums[nums[i]] {
                ans = append(ans, nums[i])
                nums[i] = -1
                break
            }
            nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
        } 
    }

    return
}