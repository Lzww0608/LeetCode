func minimumMoves(nums []int, k int, maxChanges int) int64 {
    pos := []int{}
    c := 0
    for i, x := range nums {
        if x == 0 {
            continue
        }
        pos = append(pos, i)
        c = max(c, 1)
        if i > 0 && nums[i-1] == 1 {
            if i > 1 && nums[i-2] == 1 {
                c = 3
            } else {
                c = max(c, 2)
            }
        }
    }

    c = min(c, k)
    if maxChanges >= k - c {
        return int64(max(c - 1, 0) + (k - c) * 2)
    }

    n := len(pos)
    sum := make([]int, n + 1)
    for i, x := range pos {
        sum[i+1] = sum[i] + x
    }

    ans := math.MaxInt
    sz := k - maxChanges
    for r := sz; r <= n; r++ {
        l := r - sz
        i := l + sz >> 1
        s1 := pos[i] * (i - l) - (sum[i] - sum[l])
        s2 := sum[r] - sum[i] - pos[i] * (r - i)
        ans = min(ans, s1 + s2)
    }

    return int64(ans + maxChanges * 2)
}