func countPairs(nums []int, target int) (ans int) {
    sort.Ints(nums)
    n := len(nums)
    l, r := 0, n - 1
    for l < r {
        for l < r && nums[l] + nums[r] >= target {
            r--
        }
        ans += r - l
        l++
    }

    return 
}