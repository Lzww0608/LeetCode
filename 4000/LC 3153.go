func sumDigitDifferences(nums []int) int64 {
    n := len(nums)
    ans := 0
    m := map[int]int{}
    for nums[0] > 0 {
        clear(m)
        for i, x := range nums {
            m[x%10]++
            nums[i] /= 10
        }
        sum := 0
        for _, x := range m {
            sum += x * (n - x)
        }
        ans += sum / 2
    }

    return int64(ans)
}