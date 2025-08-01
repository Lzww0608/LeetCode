func clumsy(n int) (ans int) {
    ans += solve(n)
    n -= 4
    for n > 0 {
        ans -= solve(n)
        if n >= 4 {
            ans += (n - 3) * 2
        }
        n -= 4
    }

    return
}

func solve(x int) int {
    if x == 3 {
        return 6
    } else if x <= 2 {
        return x
    }
    a, b, c, d := x, x - 1, x - 2, x - 3
    return a * b / c + d
}