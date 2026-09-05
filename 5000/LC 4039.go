const MOD int = 1_000_000_007
func sumDecoded(nums []int64) (ans int) {
    for _, t := range nums {
        w := int(t % 10)
        d := int(t / 10)
        s := strconv.Itoa(d)
        x, _ := strconv.Atoi(s[:w])
        y, _ := strconv.Atoi(s[w:])
        ans = (ans + quickPow(x, y)) % MOD
    }

    return
}

func quickPow(a, r int) int {
    res := 1
    for r > 0 {
        if r & 1 == 1 {
            res = res * a % MOD 
        }

        r >>= 1
        a = a * a % MOD
    }

    return res
}