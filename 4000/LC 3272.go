
import "math"
func countGoodIntegers(n int, k int) int64 {
    factorial := make([]int, n + 1)
    factorial[0] = 1
    for i := 1; i <= n; i++ {
        factorial[i] = factorial[i-1] * i 
    }

    ans := 0
    vis := make(map[string]bool)
    base := int(math.Pow10((n - 1) / 2))
    for i := base; i < base * 10; i++ {
        x, t := i, i 
        if n & 1 == 1 {
            t /= 10
        }
        for ; t > 0; t /= 10 {
            x = x * 10 + t % 10
        }

        if x % k != 0 {
            continue
        }

        bytes := []byte(strconv.Itoa(x))
        slices.Sort(bytes)
        s := string(bytes)
        if vis[s] {
            continue
        }
        vis[s] = true

        cnt := [10]int{}
        for _, c := range bytes {
            cnt[c - '0']++
        }

        cur := (n - cnt[0]) * factorial[n-1]
        for i := 0; i < 10; i++ {
            cur /= factorial[cnt[i]]
        }
        ans += cur
    }

    return int64(ans)
}