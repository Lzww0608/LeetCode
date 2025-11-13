func countMajoritySubarrays(nums []int, target int) (ans int) {
    n := len(nums)
    cnt := make([]int, n * 2 + 1)
    cnt[n] = 1
    cur, sum := n, 0

    for _, x := range nums {
        if x == target {
            sum += cnt[cur]
            cur++
        } else {
            cur--
            sum -= cnt[cur]
        }

        ans += sum 
        cnt[cur]++
    }

    return
}