func insert(intervals [][]int, newInterval []int) (ans [][]int) {
    a, b := newInterval[0], newInterval[1]
    n := len(intervals)
    f := true
    for i:= 0; i < n; i++ {
        v := intervals[i]
        if f && a <= v[1] {
            f = false
            if b < v[0] {
                ans = append(ans, []int{a, b})
                ans = append(ans, v)
                continue
            }
            l := min(v[0], a)
            r := max(v[1], b)
            for i + 1 < n && intervals[i+1][0] <= r {
                r = max(r, intervals[i+1][1])
                i++
            }
            ans = append(ans, []int{l, r})
        } else {
            ans = append(ans, v)
        }
        
    }

    if f {
        ans = append(ans, []int{a, b})
    }

    return 
}