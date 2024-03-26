var MOD int = int(1e9 + 7)
func countWays(ranges [][]int) int {
    sort.Slice(ranges, func(i, j int) bool {
        if ranges[i][0] == ranges[j][0] {
            return ranges[i][1] < ranges[j][1]
        }
        return ranges[i][0] < ranges[j][0]
    })

    ans := 1
    r := -1
    for _, v := range ranges {
        if v[0] > r {
            ans = (ans << 1) % MOD
        } 
        r = max(r, v[1])
    } 


    return ans
}