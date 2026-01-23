func bestTower(towers [][]int, center []int, radius int) []int {
    ans := []int{-1, -1}
    mx := -1
    for _, v := range towers {
        if abs(v[0] - center[0]) + abs(v[1] - center[1]) <= radius {
            if v[2] > mx || v[2] == mx && (v[0] < ans[0] || v[0] == ans[0] && v[1] < ans[1]) {
                mx = v[2]
                ans = v[:2]
            }
        } 
    }

    return ans
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}