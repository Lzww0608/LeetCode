func maximumRobots(chargeTimes []int, runningCosts []int, budget int64) (ans int) {
    sum, l, q := 0, 0, []int{}
    
    for r, x := range chargeTimes {
        sum += runningCosts[r]
        for len(q) > 0 && chargeTimes[q[len(q)-1]] <= x {
            q = q[:len(q)-1]
        }
        q = append(q, r)
        
        for len(q) > 0 && int64(chargeTimes[q[0]]) + int64(r - l + 1) * int64(sum) > budget {
            if l == q[0] {
                q = q[1:]
            }
            sum -= runningCosts[l]
            l++
        }
        ans = max(ans, r - l + 1)
    }

    return 
}