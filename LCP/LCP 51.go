func perfectMenu(materials []int, cookbooks [][]int, attribute [][]int, limit int) int {
    n := len(cookbooks)
    ans := 0

    m := 1 << n
    cur := make([]int, len(materials))
    for s := 1; s < m; s++ {
        copy(cur, materials)
        a, b := 0, 0
        f := true 
        for i := range n {
            if s & (1 << i) != 0 {
                a += attribute[i][0]
                b += attribute[i][1] 
                for j, x := range cookbooks[i] {
                    cur[j] -= x
                    if cur[j] < 0 {
                        f = false
                    }
                }
            }
        }
        if f && b >= limit {
            ans = max(ans, a)
        }
    }

    if ans == 0 {
        return -1
    }
    return ans
}