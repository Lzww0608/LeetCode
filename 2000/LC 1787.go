func minChanges(nums []int, k int) int {
    n := len(nums)
    m := make([]map[int]int, k)
    cnt := make([]int, k)
    for i := 0; i < k; i++ {
        m[i] = make(map[int]int)
        for j := i; j < n; j += k {
            m[i][nums[j]]++
            cnt[i]++
        }
    }

    f := make([]int, 1 << 10)
    for i := range f {
        f[i] = math.MaxInt32
    }
    f[0] = 0
    for i := 0; i < k; i++ {
        low := slices.Min(f)
        tmp := make([]int, 1 << 10)
        for j := range tmp {
            tmp[j] = low + cnt[i]
        }

        for j := 0; j < (1 << 10); j++ {
            if f[j] == math.MaxInt32 {
                continue
            }
            for p, freq := range m[i] {
                nxt := p ^ j
                tmp[nxt] = min(tmp[nxt], f[j] + cnt[i] - freq)
            }
        }
        copy(f, tmp)
    }

    return f[0]
}