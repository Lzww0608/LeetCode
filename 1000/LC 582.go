func killProcess(pid []int, ppid []int, kill int) (ans []int) {
    id := map[int]int{}
    for i, x := range pid {
        id[i] = x
    }
    m := map[int][]int{}
    for i, x := range ppid {
        m[x] = append(m[x], id[i])
    }
    //fmt.Println(m)
    q := []int{kill}
    for len(q) > 0 {
        u := q[0]
        q = q[1:]
        ans = append(ans, u)
        for _, v := range m[u] {
            q = append(q, v)
        }
    }

    return ans
}