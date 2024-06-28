func findKthNumber(k int) (ans int) {
    for i := 1; ; i++ {
        if i * quickPow(10, i) > k {
            s := strconv.Itoa(k / i)
            ans = int(s[k % i] - '0')
            break
        }
        k += quickPow(10, i)
    }

    return
}

func quickPow(x, r int) int {
    res := 1
    for ; r > 0; r >>= 1 {
        if r & 1 == 1 {
            res *= x
        }
        x *= x
    }
    return res
}
