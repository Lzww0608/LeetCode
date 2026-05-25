func minimumSwaps(nums []int) int {
    n := len(nums)
    cnt := 0
    for i, x := range nums {
        if x == 0 {
            cnt++
        }
        nums[i] = cnt
    }
    if cnt == 0 || cnt == n {
        return 0
    }
    return nums[n - cnt - 1]
}