func sortThreats(threats [][]int) [][]int {
    sort.Slice(threats, func(i, j  int) bool {
        a, b := 2 * threats[i][1] + threats[i][2], 2 * threats[j][1] + threats[j][2]
        if a == b {
            return threats[i][0] < threats[j][0]
        }

        return a > b
    })

    return threats
}