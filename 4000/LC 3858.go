func minimumOR(a [][]int) (ans int) {
    m, n := len(a), len(a[0])
    next:
    for k := 20; k >= 0; k-- {
        for i := range m {
            f := false
            for j := range n {
                x := a[i][j]
                mask := x | ans 
                if mask >> k == ans >> k {
                    f = true 
                    break
                }
            }
            if !f {
                ans |= 1 << k
                continue next
            }
        }
    }

    return ans
}