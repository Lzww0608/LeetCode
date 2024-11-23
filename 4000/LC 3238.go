func winningPlayerCount(n int, pick [][]int) (ans int) {
    m := make([]map[int]int, n)
    for i := range m {
        m[i] = make(map[int]int)
    }

    for _, v := range pick {
        a, b := v[0], v[1]
        m[a][b]++
    }
    
    for i, v := range m {
        for _, x := range v {
            if x > i {
                ans++
                break
            }
        }
    }

    return
}