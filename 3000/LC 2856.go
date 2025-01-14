func minLengthAfterRemovals(nums []int) int {
    n := len(nums)
    x := nums[n/2]
    maxCnt := sort.SearchInts(nums, x + 1) - sort.SearchInts(nums, x)

    return max(n & 1, maxCnt * 2 - n)
}