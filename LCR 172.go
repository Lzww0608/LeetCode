func countTarget(scores []int, target int) int {
    n := len(scores)
    l, r := 0, n
    for l < r {
        mid := l + ((r - l) >> 1)
        if scores[mid] > target {
            r = mid
        } else if scores[mid] < target {
            l = mid + 1 
        } else {
            i, j := mid - 1, mid + 1
            for i >= 0 && scores[i] == target {
                i--
            }
            for j < n && scores[j] == target {
                j++
            }
            return j - i - 1
        }
    }
    return 0
}