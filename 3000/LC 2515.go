func closestTarget(words []string, target string, startIndex int) int {
    n := len(words)
    ans := n + 1
    for i, word := range words {
        if word == target {
            if i >= startIndex {
                ans = min(ans, i - startIndex, n - i + startIndex)
            } else {
                ans = min(ans, startIndex - i, n - startIndex + i)
            }
        }
    }

    if ans > n {
        return -1
    }

    return ans
}