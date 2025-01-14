import "sort"
func isMajorityElement(nums []int, target int) bool {
    n := len(nums)
    maxCnt := sort.SearchInts(nums, target + 1) - sort.SearchInts(nums, target)
    return maxCnt > n / 2
}