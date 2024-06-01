func allCellsDistOrder(rows int, cols int, rCenter int, cCenter int) [][]int {
    ans := [][]int{}
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            ans = append(ans, []int{i, j})
        }
    }

    sort.Slice(ans, func(i, j int) bool {
        a := abs(ans[i][0] - rCenter) + abs(ans[i][1] - cCenter)
        b :=  abs(ans[j][0] - rCenter) + abs(ans[j][1] - cCenter)
        return a < b
    })

    return ans
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
