func longestConsecutive(nums []int) int {
    m := map[int]int{}
    for _, x := range nums {
        m[x]++
    }
    ans := 0
    for _, x := range nums {
        if m[x] == 0 {
            continue
        }
        m[x]--
        l, r := 1, 1
        for m[x-l] > 0 {
            m[x-l]--
            l++
        }
        for m[x+r] > 0 {
            m[x+r]--
            r++
        }
        ans = max(ans, l + r - 1)
    }
    return ans
}