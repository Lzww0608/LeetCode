func findLatestStep(arr []int, m int) int {
    n := len(arr)
    if n == m {
        return n
    }

    length := make([]int, n + 2)
    cnt := make([]int, n + 1)
    ans := -1
    for i, x := range arr {
        l, r := length[x-1], length[x+1]
        cur := l + r + 1
        length[x-l] = cur
        length[x+r] = cur
        cnt[cur]++
        cnt[l]--
        cnt[r]--
        if cnt[m] > 0 {
            ans = i + 1
        }
    }

    return ans
}