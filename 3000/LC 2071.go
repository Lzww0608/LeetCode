func maxTaskAssign(tasks []int, workers []int, pills int, strength int) (ans int) {
    n, m := len(tasks), len(workers)
    sort.Ints(tasks)
    sort.Ints(workers)

    check := func(target int) bool {
        a := tasks[:target]
        b := workers[m-target:]
        cnt := 0
        q := []int{}
        j := target - 1
        for i := target - 1; i >= 0 && cnt <= pills; i-- {
            x := a[i]
            for j >= 0 && b[j] + strength >= x {
                q = append(q, b[j])
                j--
            }
            
            if len(q) == 0 {
                return false
            } else if q[0] >= x {
                q = q[1:]
            } else {
                cnt++
                q = q[:len(q)-1]
            }
            
        }

        return cnt <= pills
    } 

    l, r := 0, min(m, n) + 1
    for l < r {
        mid := l + ((r - l) >> 1)
        if check(mid) {
            ans = mid 
            l = mid + 1
        } else {
            r = mid
        }
    }

    return 
}