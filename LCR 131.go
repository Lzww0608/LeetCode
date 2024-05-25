func cuttingBamboo(n int) (ans int) {
    if n < 4 {
        return 2
    }

    if n % 3 == 0 {
        ans = quickPow(3, n / 3)
    } else if n % 3 == 1 {
        ans = quickPow(3, n / 3 - 1) * 4
    } else {
        ans = quickPow(3, n / 3) * 2
    }

    return ans
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