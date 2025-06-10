const N int = 26
func maxDifference(s string) int {
    cnt := [N]int{}
    for _, c := range s {
        x := int(c - 'a')
        cnt[x]++
    }

    a := [2]int{math.MaxInt32, math.MinInt32}
    b := [2]int{math.MaxInt32, math.MinInt32}
    for _, x := range cnt {
        if x == 0 {
            continue
        }
        if x & 1 == 1 {
            a[0] = min(a[0], x)
            a[1] = max(a[1], x)
        } else {
            b[0] = min(b[0], x)
            b[1] = max(b[1], x)
        }
    }

    return a[1] - b[0]
}