func longestConsecutive(nums []int) (ans int) {
    m := map[int]bool{}
    for _, x := range nums {
        m[x] = true
    }
    for _, x := range nums {
        if !m[x] {
            continue
        }
        l, r := x - 1, x + 1
        m[x] = false
        for m[l] {
            m[l] = false
            l--
        }
        for m[r] {
            m[r] = false
            r++
        }
        ans = max(ans, r - l - 1)
    }

    return
}