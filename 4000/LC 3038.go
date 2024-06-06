func maxOperations(nums []int) (ans int) {
    score := nums[0] + nums[1]
    for i := 2; i < len(nums); i += 2 {
        if i + 1 < len(nums) && nums[i] + nums[i + 1] == score {
            ans++
        } else {
            break
        }
    }

    return ans + 1
}