func canReach(arr []int, start int) bool {
    q := []int{start}
    n := len(arr)
    vis := make([]bool, n)
    
    for len(q) > 0 {
        t := q[0]
        q = q[1:]
        vis[t] = true
        if arr[t] == 0 {
            return true
        } 
        if t + arr[t] < n && ! vis[t+arr[t]] {
            q = append(q, t + arr[t])
        }
        if t - arr[t] >= 0 && !vis[t-arr[t]] {
            q = append(q, t - arr[t])
        }
    }

    return false
}