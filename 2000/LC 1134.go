func isArmstrong(n int) bool {
    a := []int{}
    for t := n; t > 0; t /= 10 {
        a = append(a, t % 10)
    }

    sum := 0
    for _, x := range a {
        sum += quickPow(x, len(a))
    }

    return sum == n
}


func quickPow(a, r int) int {
    res := 1
    for r > 0 {
        if r & 1 == 1 {
            res *= a
        }
        a *= a
        r >>= 1
    }
    return res
}