func maxStudentsOnBench(students [][]int) (ans int) {
    cnt := make(map[int]int)
    vis := make(map[[2]int]bool)

    for _, v := range students {
        if vis[[2]int{v[0], v[1]}] {
            continue
        }
        vis[[2]int{v[0], v[1]}] = true 
        cnt[v[1]]++
        ans = max(ans, cnt[v[1]])
    }

    return
}