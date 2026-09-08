func countGoodRotations(nums []int) (ans int) {
    n := len(nums)
    nums = append(nums, nums...)
    a, b := 0, 0
    for i, x := range nums[:n] {
        if i < n / 2 {
            a += x
        } else {
            b += x
        }
    }

    if a > b {
        ans++
    }

    for i := n; i < n * 2 - 1; i++ {
        b += nums[i] - nums[i - n / 2]
        a += nums[i - n / 2] - nums[i - n]
        if a > b {
            ans++
        }
    }

    return 
}