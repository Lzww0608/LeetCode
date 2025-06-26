func timeTaken(arrival []int, state []int) []int {
    n := len(arrival)
    ans := make([]int, n)
    door, t := 1, 0
    cur := [2]int{}

    for min(cur[0], cur[1]) < n {
        if arrival[min(cur[0], cur[1])] > t {
            t = arrival[min(cur[0], cur[1])]
            door = 1
        }
        
        for cur[door] < n && arrival[cur[door]] <= t {
            if state[cur[door]] == door {
                ans[cur[door]] = t
                t++
            }
            cur[door]++
        }
        door ^= 1
    }

    return ans
}