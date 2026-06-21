func minLights(lights []int) (ans int) {
    n := len(lights)
    dif := make([]int, n + 1)
    for i, x := range lights {
        if x == 0 {
            continue
        }
        l := max(0, i - x)
        r := min(n, i + x + 1)
        dif[l]++
        dif[r]--
    }

    cur := 0
    for i := range dif {
        cur += dif[i]
        dif[i] = cur 
    }

    cur = 0
    for i := range n {
        if dif[i] != 0 {
            ans += (cur + 2) / 3
            cur = 0
        } else {
            cur++
        }
    }

    ans += (cur + 2) / 3
    return
}