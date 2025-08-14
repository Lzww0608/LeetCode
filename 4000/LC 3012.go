func minimumArrayLength(nums []int) int {
    mn := slices.Min(nums)
    ans := 0
    for _, x := range nums {
        if x % mn != 0 {
            return 1
        }
        if x == mn {
            ans++
        }
    }

    return (ans + 1) / 2
}