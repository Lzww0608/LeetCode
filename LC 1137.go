func tribonacci(n int) int {
    if n == 0 {
        return 0
    } else if n == 1 || n == 2 {
        return 1
    }

    mat := [3][3]int{[3]int{1, 1, 1}, [3]int{1, 0, 0}, [3]int{0, 1, 0}}
    mul := func(a, b [3][3]int) [3][3]int {
        t := [3][3]int{}

        for i := 0; i < 3; i++ {
            for j := 0; j < 3; j++ {
                t[i][j] = a[i][0] * b[0][j] + a[i][1] * b[1][j] + a[i][2] * b[2][j]
            }
        }

        return t
    }

    ans := [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
    for k := n; k > 0; k >>= 1 {
        if k & 1 == 1 {
            ans = mul(ans, mat)
        }
        mat = mul(mat, mat)
    }
    return ans[0][2]
}