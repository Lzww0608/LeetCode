const N int = 10
func score(cards []string, c byte) (ans int) {
    pre := [N]int{}
    suf := [N]int{}
    sum1, sum2 := 0, 0
    mx1, mx2 := 0, 0

    add := 0
    k := int(c - 'a')
    for _, card := range cards {
        x, y := int(card[0] - 'a'), int(card[1] - 'a')
        if x == k && y == k {
            add++
        } else if x == k {
            sum1++
            pre[y]++
            mx1 = max(mx1, pre[y])
        } else if y == k {
            sum2++
            suf[x]++
            mx2 = max(mx2, suf[x])
        }
    }

    if add >= sum1 + sum2 {
        return sum1 + sum2
    }

    if sum1 > 0 {
        x := sum1 - mx1
        y := min(sum1 / 2, x)
        ans += y 
        ans += min(add, max(0, sum1 - y * 2))
        add -= min(add, max(0, sum1 - y * 2))
    }
    
    if sum2 > 0 {
        x := sum2 - mx2
        y := min(sum2 / 2, x)
        ans += y 
        ans += min(add, max(0, sum2 - y * 2))
        add -= min(add, max(0, sum2 - y * 2))
    }

    ans += add / 2
    return  
}