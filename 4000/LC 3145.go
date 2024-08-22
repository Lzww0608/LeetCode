func sumE(k int) (res int) {
    var n, cnt, sum int
    for i := bits.Len(uint(k + 1)) - 1; i > 0; i-- {
        c := cnt << i + i << (i - 1)
        if c <= k {
            k -= c
            res += sum << i + i * (i - 1) / 2 << (i - 1)
            sum += i
            cnt++
            n |= 1 << i
        }
    }

    if cnt <= k {
        k -= cnt
        res += sum
        n |= 1
    }

    for ; k > 0; k-- {
        res += bits.TrailingZeros(uint(n))
        n &= n - 1
    }

    return
}

func findProductsOfElements(queries [][]int64) []int {
    ans := make([]int, len(queries))
    for i, q := range queries {
        r := sumE(int(q[1]) + 1)
        l := sumE(int(q[0]))
        ans[i] = quickPow(2, r - l, int(q[2]))
    }

    return ans
}


func quickPow(a, r, mod int) int {
    res := 1 % mod
    for r > 0 {
        if r & 1 == 1 {
            res = (res * a) % mod
        }
        a = a * a % mod
        r >>= 1
    }

    return res
}