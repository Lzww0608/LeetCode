func maxTotalFruits(fruits [][]int, startPos int, k int) (ans int) {
    n := len(fruits)
    l := sort.Search(n, func(i int) bool {
        return fruits[i][0] >= startPos - k
    })

    for r, s := l, 0; r < n && fruits[r][0] <= startPos + k; r++ {
        s += fruits[r][1]
        for fruits[r][0] * 2 - fruits[l][0] - startPos > k && fruits[r][0] - fruits[l][0] * 2 + startPos > k {
            s -= fruits[l][1]
            l++
        }

        ans = max(ans, s)
    }

    return
}