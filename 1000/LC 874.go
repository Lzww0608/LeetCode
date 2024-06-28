var dir [4][2]int = [4][2]int{[2]int{0, 1}, [2]int{1, 0}, [2]int{0, -1}, [2]int{-1, 0}}
func robotSim(commands []int, obstacles [][]int) int {
    x, y, twd := 0, 0, 0
    ans := 0
    m := map[int]bool{}
    for _, x := range obstacles {
        m[x[0]*60001 + x[1]] = true
    }
    for _, command := range commands {
        if command < 0 {
            if command == -1 {
                twd = (twd + 1) % 4
            } else {
                twd = (twd + 3) % 4
            }
            continue
        }
        for i := 0; i < command; i++ {
            x, y = x + dir[twd][0], y + dir[twd][1]
            if m[x * 60001 + y] == true {
                x, y = x - dir[twd][0], y - dir[twd][1]
                break
            }
        }
        ans = max(ans, x * x + y * y)
    }
    return ans 
}
