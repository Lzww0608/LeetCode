func minOperations(boxes string) []int {
    n := len(boxes)
    suf := make([]int, n + 1)
    cnt := 0
    for i := n - 1; i >= 0; i-- {
        suf[i] = suf[i + 1] + cnt
        if boxes[i] == '1' {
            cnt++
        }
    }

    ans := make([]int, n)
    pre, cnt := 0, 0
    for i := range boxes {
        pre += cnt
        ans[i] = pre + suf[i]
        if boxes[i] == '1' {
            cnt++
        }
    }

    return ans
}