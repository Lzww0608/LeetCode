func maxSubarrays(nums []int) (ans int) {
    sum := (1 << 30) - 1
    for _, x := range nums {
        sum &= x
        if sum == 0 {
            ans++
            sum = (1 << 30) - 1
        }
    }

    return max(1, ans)
}