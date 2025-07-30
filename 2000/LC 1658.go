func minOperations(nums []int, x int) int {
    sum := 0
    for _, y := range nums {
        sum += y 
    }
    if sum < x {
        return -1
    }

    ans := -1
    cur := 0
    for l, r := 0, 0; r < len(nums); r++ {
        cur += nums[r]
        for cur > sum - x {
            cur -= nums[l]
            l++
        }
        if cur == sum - x {
            ans = max(ans, r - l + 1)
        }
    }

    if ans == -1 {
        return -1
    }
    return len(nums) - ans
}