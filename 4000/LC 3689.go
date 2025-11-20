func maxTotalValue(nums []int, k int) int64 {
    mx := slices.Max(nums)
    mn := slices.Min(nums)

    return int64(k * (mx - mn))
}