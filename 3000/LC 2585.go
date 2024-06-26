const MOD int = 1e9 + 7 
const N int = 1005
func waysToReachTarget(target int, types [][]int) int {
    f := make([]int, target + 1)
    f[0] = 1

    for _, t := range types {
        k, v := t[0], t[1]
        for s := target; s > 0; s-- {
            for j := 1; j <= k && s - j * v >= 0; j++ {
                f[s] += f[s - j * v]  
                f[s] %= MOD
            }
        }
    }

    return f[target]
}
