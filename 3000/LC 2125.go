func numberOfBeams(bank []string) (ans int) {
    pre := 0
    for _, row := range bank {
        cur := 0
        for _, c := range row {
            if c == '1' {
                cur++
            } 
        }
        if cur > 0 {
            ans += cur * pre
            pre = cur
        }
    }

    return
}