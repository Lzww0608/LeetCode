func minTime(skill []int, mana []int) int64 {
    n := len(skill)
    f := make([]int, n)
    pre := 0
    for i := range n {
        pre += skill[i] * mana[0]
        f[i] = pre
    }
    
    fmt.Println(f)
    for _, x := range mana[1:] {
        mx := 0
        pre = f[0]
        for i := range n {
            mx = max(mx, f[i] - pre)
            pre += skill[i] * x 
        }

        pre = 0
        cur := mx + f[0]
        for i := range n {
            pre += skill[i] * x
            f[i] = cur + pre 
        }

        //fmt.Println(cur, f)
    }

    return int64(f[n - 1])
}