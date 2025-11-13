func countMajoritySubarrays(nums []int, target int) int64 {
    ans := 0
    n := len(nums) * 2 + 2
    f := make([]int, n)

    update := func(i int) {
        for ; i < n; i += i & (-i) {
            f[i] += 1
        }
    }

    query := func(i int) (res int) {
        for ; i > 0; i -= i & (-i) {
            res += f[i]
        }

        return
    }

    cnt := len(nums) + 1
    update(cnt)
    for _, x := range nums {
        if x == target {
            cnt++
        } else {
            cnt--
        }
        
        ans += query(cnt - 1)
        update(cnt)
    }

    return int64(ans)
}