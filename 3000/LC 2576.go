import (
    "sort"
)
func maxNumOfMarkedIndices(nums []int) int {
    sort.Ints(nums)
    n := len(nums)

    l := 0
    r := (n + 1) / 2
    for r < n {
        for r < n && nums[r] < 2 * nums[l] {
            r++
        } 
        if r < n {
            l++
        }
        r++
    }

    return l * 2
}