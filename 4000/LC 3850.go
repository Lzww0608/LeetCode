func countSequences(nums []int, k int64) int {
    n := len(nums)
    m := make([]map[pair]int, n)
    for i := range m {
        m[i] = make(map[pair]int)
    }

    var dfs func(int, int, int) int 
    dfs = func(i, a, b int) int {
        g := gcd(a, b)
        a /= g 
        b /= g
        if i == n {
            if a == int(k) * b {
                return 1
            }
            return 0
        }
        if v, ok := m[i][pair{a, b}]; ok {
            return v
        }

        // d := n - i 
        // if a * quickPow(6, d) < int(k) * b {
        //     return 0
        // }
        // if int(k) * quickPow(6, d) * b < a {
        //     return 0
        // }
        
        res := dfs(i + 1, a, b) + dfs(i + 1, a * nums[i], b) + dfs(i + 1, a, b * nums[i])
        m[i][pair{a, b}] = res 
        return res
    }

    return dfs(0, 1, 1)
}

func quickPow(a, r int) int {
    res := 1
    for r > 0 {
        if r & 1 == 1 {
            res = res * a
        }
        a = a * a
        r >>= 1
    }
    return res
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a % b
    }
    return a
}

type pair struct {
    x, y int
}