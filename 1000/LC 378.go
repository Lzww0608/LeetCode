func kthSmallest(matrix [][]int, k int) int {
    m, n := len(matrix), len(matrix[0])
    l, r := matrix[0][0], matrix[m-1][n-1]

    check := func(target int) bool {
        cnt := 0
        r, c := n - 1, 0
        for r >= 0 && c < n {
            if target >= matrix[r][c] {
                cnt += r + 1
                c++
            } else {
                r--
            }
        }

        return cnt >= k
    }

    for l < r {
        mid := l + ((r - l) >> 1)
        if check(mid) {
            r = mid
        } else {
            l = mid + 1
        }
    }

    return l
}