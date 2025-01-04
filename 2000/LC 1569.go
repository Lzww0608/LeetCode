const N int = 1001
const MOD int = 1_000_000_007
var factor [N * 2]int
var rev [N * 2]int

func init() {
    factor[0], rev[0] = 1, 1
    cur := 1
    for i := 1; i < len(factor); i++ {
        cur = (cur * i) % MOD
        factor[i] = cur
        rev[i] = modInverse(factor[i])
    }
}

func modInverse(x int) int {
    return quickPow(x, MOD - 2) 
}

func quickPow(a, r int) int {
    if a == 1 {
        return 1
    }
    res := 1
    for r > 0 {
        if r & 1 == 1 {
            res = (res * a) % MOD
        }
        r >>= 1
        a = (a * a) % MOD
    }

    return res
}

func numOfWays(nums []int) (ans int) {
    return dfs(nums) - 1
}

func dfs(a []int) (res int) {
    if len(a) <= 1 {
        return 1
    }

    l := []int{}
    r := []int{}
    for _, x := range a[1:] {
        if x > a[0] {
            r = append(r, x)
        } else if x < a[0] {
            l = append(l, x)
        }
    }

    res = rev[len(l)] % MOD * factor[len(l) + len(r)] % MOD * rev[len(r)] % MOD
    res = res * dfs(l) % MOD * dfs(r) % MOD
    return
}