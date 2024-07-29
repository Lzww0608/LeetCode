func getGoodIndices(variables [][]int, target int) (ans []int) {
    

    for i, v := range variables {
        if quickPow((quickPow(v[0], v[1], 10) % 10), v[2], v[3]) == target {
            ans = append(ans, i)
        }
    }

    return
}

func quickPow(a, r, x int) int {
    res := 1
    for ; r > 0; r >>= 1 {
        if r & 1 == 1 {
            res *= a
            res %= x
        }
        a *= a
        a %= x
    }

    return res
}