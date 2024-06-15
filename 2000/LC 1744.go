func canEat(candiesCount []int, queries [][]int) []bool {
    n, m := len(queries), len(candiesCount)
    ans, prefix := make([]bool, n), make([]int, m+1)
    prefix[0] = 0
    for i, x := range candiesCount {
        prefix[i+1] = prefix[i] + x
    } 
    for i, x := range queries {
        a, b, c := x[0], x[1], x[2]
        if prefix[a+1] > b && prefix[a] < c * (b + 1) {
            ans[i] = true
        } else {
            ans[i] = false
        } 
    }
    return ans

}
