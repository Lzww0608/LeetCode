func leastBricks(wall [][]int) int {

    cnt := make(map[int]int)
    for _, v := range wall {
        cur := 0 
        for _, x := range v[:len(v)-1] {
            cur += x 
            cnt[cur]++
        }
    }

    n := len(wall)
    ans := n 
    for _, x := range cnt {
        ans = min(ans, n - x)
    }

    return ans
}