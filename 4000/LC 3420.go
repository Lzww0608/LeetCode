func countNonDecreasingSubarrays(nums []int, k int) int64 {
    type pair struct {
        x, sz int
    }
    n := len(nums)
    ans, cnt := 0, 0
    q := []pair{}

    for l, r := n - 1, n - 1; l >= 0; l-- {
        x := nums[l]
        sz := 1
        for len(q) > 0 && x >= q[len(q) - 1].x {
            cur := q[len(q) - 1]
            q = q[:len(q) - 1]
            sz += cur.sz
            cnt += (x - cur.x) * cur.sz
        }
        q = append(q, pair{x, sz})

        for cnt > k {
            cnt -= q[0].x - nums[r]
            r--
            q[0].sz--
            if q[0].sz == 0 {
                q = q[1:]
            }
        }

        ans += r - l + 1
    }

    return int64(ans)
}