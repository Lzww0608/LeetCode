func isMiddleElementUnique(nums []int) bool {
    n := len(nums)
    cnt := 0
    for _, x := range nums {
        if x == nums[n / 2] {
            cnt++
        }
    }

    return cnt == 1
}