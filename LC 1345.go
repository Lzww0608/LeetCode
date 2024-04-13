func minJumps(arr []int) int {
    q := []int{0}
    n := len(arr)
    vis := make([]bool, n)
    m := map[int][]int{}
    for i, x := range arr {
        m[x] = append(m[x], i)
    }

    k := 0
    for len(q) > 0 {
        for i := len(q); i > 0; i-- {
            t := q[0]
            q = q[1:]
            vis[t] = true
            if t == n - 1 {
                return k
            }
            if t + 1 < n && !vis[t+1] {
                q = append(q, t + 1)
            }
            if t - 1 >= 0 && !vis[t-1] {
                q = append(q, t - 1)
            }
            for _, x := range m[arr[t]] {
                if !vis[x] {
                    q = append(q, x)
                } 
            } 
            delete(m, arr[t])
        }
        k++
    }

    return -1
}