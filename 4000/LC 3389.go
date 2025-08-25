
import "math"
const N int = 26
func makeStringGood(s string) int {
    cnt := make([]int, N)
    for i := range s {
        cnt[int(s[i] - 'a')]++
    }

    mx := slices.Max(cnt)
    f := make([]int, N + 1)
    ans := math.MaxInt
    for target := 1; target <= mx; target++ {
        f[N - 1] = min(cnt[N - 1], abs(target - cnt[N - 1]))
        for i := N - 2; i >= 0; i-- {
            x, y := cnt[i], cnt[i + 1]
            f[i] = f[i + 1] + min(x, abs(x - target))

            if y < target {
                t := 0
                if x > target {
                    t = target
                }

                f[i] = min(f[i], f[i + 2] + max(x - t, target - y))
            }
        }

        ans = min(ans, f[0])
    }

    return ans
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}