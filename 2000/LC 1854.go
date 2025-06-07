const MIN int = 1950
const MAX int = 2050
func maximumPopulation(logs [][]int) int {
    dif := make([]int, MAX - MIN + 1)
    for _, log := range logs {
        dif[log[0] - MIN]++
        dif[log[1] - MIN]--
    }

    sum := 0
    ans, mx := 0, 0
    for i := range dif {
        sum += dif[i]
        if sum > mx {
            mx = sum 
            ans = i + MIN
        }
    }

    return ans
}