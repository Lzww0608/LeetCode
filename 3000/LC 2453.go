func destroyTargets(nums []int, space int) (ans int) {
    m := make(map[int]int)
    for _, x := range nums {
        t := x % space
        if _, ok := m[t]; ok {
            m[t] = min(m[t], x)
        } else {
            m[t] = x
        }
    }
    cnt := make(map[int]int)
    for _, x := range nums {
        x = m[x%space]
        cnt[x]++
        if cnt[x] > cnt[ans] || cnt[x] == cnt[ans] && ans > x {
            ans = x
        }
    }

    return

    
}