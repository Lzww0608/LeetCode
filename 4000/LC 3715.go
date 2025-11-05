const N int = 100_001

var (
    isPrimes [N]bool 
    f [N]int
)

func init() {
    for i := range N {
        f[i] = 1
    }

    for i := 2; i < N; i++ {
        if isPrimes[i] {
            continue
        }

        for j := i; j < N; j += i {
            isPrimes[j] = true 
            cnt := 0
            for x := j; x % i == 0; x /= i {
                cnt++
            }
            if cnt & 1 == 1 {
                f[j] *= i
            }
        }
    }
}

func sumOfAncestors(n int, edges [][]int, nums []int) int64 {
    g := make([][]int, n)
    for _, edge := range edges {
        u, v := edge[0], edge[1]
        g[u] = append(g[u], v)
        g[v] = append(g[v], u)
    }

    ans := 0
    cnt := make(map[int]int)
    var dfs func(int, int)
    dfs = func(u, fa int) {
        ans += cnt[f[nums[u]]]
        cnt[f[nums[u]]]++
        for _, v := range g[u] {
            if v == fa {
                continue
            }
        
            dfs(v, u)
        }
        cnt[f[nums[u]]]--
    }
    dfs(0, -1)

    return int64(ans)
}
