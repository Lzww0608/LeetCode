func largestNumber(cost []int, target int) string {
    f := make([]int, target + 1)
    for i := range f {
        f[i] = math.MinInt32
    }
    f[0] = 0

    for _, c := range cost {
        for j := c; j <= target; j++ {
            f[j] = max(f[j], f[j-c] + 1)
        }
    }

    if f[target] < 0 {
        return "0"
    }

    ans := []byte{}
    for i, j := 8, target; i >= 0; i-- {
        for c := cost[i]; j >= c && f[j] == f[j-c] + 1; j -= c {
            ans = append(ans, byte('1' + i))
        }
    }

    return string(ans)

}