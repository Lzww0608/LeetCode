func findTargetIn2DPlants(plants [][]int, target int) bool {
    if len(plants) == 0 {
        return false
    }
    m, n := len(plants), len(plants[0])
    i, j := 0, n - 1
    for i < m && j >= 0 {
        if plants[i][j] > target {
            j--
        } else if plants[i][j] < target {
            i++
        } else {
            return true
        }
    }

    return false
}