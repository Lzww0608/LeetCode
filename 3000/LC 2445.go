func numberOfNodes(n int, queries []int) (ans int) {
    m := make(map[int]int)
    for _, x := range queries {
        m[x]++
    }

    var dfs func(int, int) 
    dfs = func(x, cnt int) {
        if x > n {
            return
        }
        if m[x] & 1 == 1{
            cnt++
        }
        if cnt & 1 == 1 {
            ans++
        }
        dfs(x * 2, cnt)
        dfs(x * 2 + 1, cnt)
    }
    dfs(1, 0)

    return
}