
func maximizeSquareArea(m int, n int, hFences []int, vFences []int) int {
    ans := -1
    cnt := make(map[int]bool)
    hFences = append(hFences, 1)
    hFences = append(hFences, m)
    sort.Ints(hFences)

    for i := 0; i < len(hFences); i++ {
        for j := i + 1; j < len(hFences); j++ {
            cnt[hFences[j] - hFences[i]] = true
        }
    }

    vFences = append(vFences, 1)
    vFences = append(vFences, n)
    sort.Ints(vFences)
func maximizeSquareArea(m int, n int, hFences []int, vFences []int) int {
    ans := -1
    cnt := make(map[int]bool)
    hFences = append(hFences, 1)
    hFences = append(hFences, m)
    sort.Ints(hFences)

    for i := 0; i < len(hFences); i++ {
        for j := i + 1; j < len(hFences); j++ {
            cnt[hFences[j] - hFences[i]] = true
        }
    }

    vFences = append(vFences, 1)
    vFences = append(vFences, n)
    sort.Ints(vFences)
    for i := 0; i < len(vFences); i++ {
        for j := i + 1; j < len(vFences); j++ {
            d := vFences[j] - vFences[i]
            if cnt[d] {
                ans = max(ans, d * d)
            }
        }
    }

    return ans % 1_000_000_007
}
    for i := 0; i < len(vFences); i++ {
        for j := i + 1; j < len(vFences); j++ {
            d := vFences[j] - vFences[i]
            if cnt[d] {
                ans = max(ans, d * d)
            }
        }
    }

    return ans % 1_000_000_007
}