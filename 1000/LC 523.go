func checkSubarraySum(nums []int, k int) bool {
    cnt := make(map[int]bool)
    pre := 0

    for _, x := range nums {
        cur := (pre + x) % k
        if cnt[cur] {
            return true
        }
        cnt[pre] = true 
        pre = cur
    }

    return false
}