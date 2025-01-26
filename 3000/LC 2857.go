func countPairs(coordinates [][]int, k int) (ans int) {
    cnt := make(map[[2]int]int)
    for _, v := range coordinates {
        x1, y1 := v[0], v[1]
        for a := 0; a <= k; a++ {
            b := k - a
            x2, y2 := a ^ x1, b ^ y1 
            ans += cnt[[2]int{x2, y2}]
        }
        cnt[[2]int{x1, y1}]++
    }

    return
}