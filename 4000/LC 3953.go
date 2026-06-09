const N int = 100_001

var F [N][]int 

func init() {
    for i := 2; i < N; i++ {
        if len(F[i]) != 0 {
            continue
        }

        for j := i; j < N; j += i {
            F[j] = append(F[j], i)
        }
    }
}

func maxScore(nums []int, maxVal int) int {
    mx := maxVal
    f, g := [N]int{}, [N]int{}
    vis := make(map[int]bool)
    for _, x := range nums {
        f[x]++
        vis[x] = true
        mx = max(mx, x)
    }

    for i := 1; i <= mx; i++ {
        for j := i * 2; j <= mx; j += i {
            f[i] += f[j]
        }
    }

    for i := 2; i <= mx; i++ {
        d := len(F[i])
        for j := 1; j < (1 << d); j++ {
            t := 1
            for k := range d {
                if j & (1 << k) != 0 {
                    t *= F[i][k]
                }
            }
            cnt := bits.OnesCount(uint(j))
            if cnt & 1 == 1 {
                g[i] += f[t]
            } else {
                g[i] -= f[t]
            }
        }
    }

    ans := 0
    if vis[1] {
        ans = 1
    }
    for i := 2; i <= mx; i++ {
        if vis[i] {
            ans = max(ans, i - (g[i] - 1))
        } else if i <= maxVal {
            ans = max(ans, i - max(1, g[i]))
        }
    }

    return ans
}