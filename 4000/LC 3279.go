func maxArea(height int, positions []int, directions string) int64 {
    period := height * 2
    changes := make(map[int]int)

    solve := func(start, dir int) {
        a := height - start 
        if dir < 0 {
            a = start 
        }
        b := height + a 
        changes[a] -= 2 * dir
        changes[b] += 2 * dir
    }

    sum, dif := 0, 0
    for i, x := range positions {
        dir := 1
        if directions[i] == 'D' {
            dir = -1
        }
        sum += x
        dif += dir
        solve(x, dir)
    }

    ans := sum 
    for i := 1; i < period; i++ {
        sum += dif 
        dif += changes[i]
        ans = max(ans, sum)
    }

    return int64(ans)
}