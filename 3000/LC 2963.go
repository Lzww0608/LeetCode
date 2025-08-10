const MOD int = 1_000_000_007
func numberOfGoodPartitions(nums []int) int {
    l := make(map[int]int)
    r := make(map[int]int)

    for i, x := range nums {
        if _, ok := l[x]; !ok {
            l[x] = i
            r[x] = i
        } else {
            r[x] = i
        }
    }

    d := 0
    right := -1
    for _, x := range nums {
        if l[x] > right {
            d++
            right = r[x]
        } else {
            right = max(right, r[x])
        }
    }

    return quickPow(2, d - 1)
}

func quickPow(a, r int) int {
    res := 1 
    for r > 0 {
        if r & 1 == 1 {
            res = res * a % MOD
        }
        a = a * a % MOD
        r >>= 1
    }

    return res
}