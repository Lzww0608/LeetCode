func maxAliveYear(birth []int, death []int) (ans int) {
    dif := make([]int, 102)
    mx := 0
    for i := range birth {
        x, y := birth[i] - 1900, death[i] - 1900 + 1
        dif[x]++
        dif[y]--
    }

    sum := 0
    for i, x := range dif {
        sum += x
        if sum > mx {
            ans, mx = i + 1900, sum
        }
    }

    return
}