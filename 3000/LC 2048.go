func nextBeautifulNumber(n int) int {
    n++
    for {
        cnt := [10]int{}
        for t := n; t > 0; t /= 10 {
            cnt[t % 10]++
        }
        f := true
        for k, v := range cnt {
            if v != 0 && v != k {
                f = false
                break
            }
        }

        if f {
            break
        }
        n++
    }

    return n
}