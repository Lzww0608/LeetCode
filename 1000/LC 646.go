func findLongestChain(pairs [][]int) (ans int) {
    sort.Slice(pairs, func(i, j int) bool {
        if pairs[i][0] == pairs[j][0] {
            return pairs[i][1] < pairs[j][1]
        }
        return pairs[i][0] < pairs[j][0]
    })

    pre := math.MinInt32
    for _, pair := range pairs {
        l, r := pair[0], pair[1]
        if l > pre {
            ans++
            pre = r 
        } else if r < pre {
            pre = r
        }
    }

    return
}