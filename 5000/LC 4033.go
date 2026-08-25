func validSubarrays(nums []int, k int, queries [][]int) []bool {
    n := len(nums)
    m := make(map[int]int)
    for _, x := range nums {
        if _, ok := m[x]; !ok {
            m[x] = rand.Intn(math.MaxInt)
        }
    }

    pre := make([]int, n + 1)
    for i, x := range nums {
        pre[i + 1] = pre[i] ^ m[x]
    }

    solve := func(t int) []int {
        p := make([]int, n)
        cnt := make(map[int]int)
        for l, r := 0, 0; r < n; r++ {
            cnt[nums[r]]++
            for len(cnt) >= t {
                if cnt[nums[l]]--; cnt[nums[l]] == 0 {
                    delete(cnt, nums[l])
                }
                l++
            }

            p[r] = l
        }

        return p
    }

    l1 := solve(k)
    l2 := solve(k + 1)
    ans := make([]bool, len(queries))
    for i, q := range queries {
        l, r := q[0], q[1]
        ans[i] = (pre[r + 1] ^ pre[l] == 0) && l >= l2[r] && l < l1[r]
    }

    return ans
}