func firstUniqChar(s string) int {
    m := map[byte]int{}

    for i := range s {
        if _, ok := m[s[i]]; !ok {
            m[s[i]] = i
        } else {
            m[s[i]] = math.MaxInt32
        }
    }

    ans := math.MaxInt32
    for _, v := range m {
        ans = min(v, ans)
    }

    if ans == math.MaxInt32 {
        return -1
    }
    return ans
}
