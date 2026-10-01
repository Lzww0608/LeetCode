func longestSubarray(nums []int, k int) (ans int) {
    first := make([]int, k)
    last := make([]int, k)
    for i := 1; i < k; i++ {
        first[i], last[i] = -1, -1
    }

    sum := 0
    p := make([][]int, k)
    for i, x := range nums {
        sum = ((sum + x) % k + k) % k
        if first[sum] == -1 {
            first[sum] = i + 1
        } else {
            ans = max(ans, i + 1 - first[sum])
        }
        last[sum] = i + 1
        p[(x * 2 % k + k) % k] = append(p[(x * 2 % k + k) % k], i)
    }

    for i, x := range first {
        if x < 0 {
            continue
        }
        for j, y := range last {
            if y - x <= ans {
                continue
            }

            v := p[((j - i) % k + k) % k]
            idx := sort.SearchInts(v, x)
            if idx < len(v) && v[idx] < y {
                ans = y - x
            }
        }
    }

    return
}