var f = make([]int, 20)

func numTrees(n int) int { 
    if n == 0 {
        return 1
    } else if n < 3 {
        return n
    } else if f[n] != 0 {
        return f[n]
    }

    ans := 0
    for i := 1; i <= n; i++ {
        l, r := numTrees(i - 1), numTrees(n - i)
        ans += l * r
    }

    f[n] = ans

    return ans
}