func queensAttacktheKing(queens [][]int, king []int) [][]int {
    x, y := king[0], king[1]
    ans := [][]int{}
    m := map[int]bool{}
    for _, q := range queens {
        m[q[0]*8+q[1]] = true
    }
    find := func(a, b int) []int {
        i, j := x + a, y + b
        for i >= 0 && i < 8 && j >= 0 && j < 8 {
            if m[i*8+j] {
                return []int{i,j} 
            }
            i, j = i + a, j + b
        }
        return []int{-1}
    }
    dir := [][2]int{{0,1}, {0,-1}, {1,0}, {-1,0}, {1,1}, {1,-1}, {-1,1}, {-1,-1}}
    for i := 0; i < 8; i++ {
        t := find(dir[i][0], dir[i][1])
        if t[0] != -1 {
            ans = append(ans, t)
        }
    }
    return ans
}