func vowelStrings(words []string, queries [][]int) []int {
    m := map[byte]bool{
        'a': true,
        'e': true,
        'i': true,
        'o': true,
        'u': true,
    }

    f := make([]int, len(words) + 1)
    for i, s := range words {
        if m[s[0]] && m[s[len(s)-1]] {
            f[i + 1] = 1
        } else {
            f[i + 1] = 0
        }
    }

    for i := 0; i < len(words); i++ {
        f[i + 1] += f[i]
    }

    ans := make([]int, len(queries))
    for i, q := range queries {
        l, r := q[0], q[1]

        ans[i] = f[r + 1] - f[l]
    }

    return ans
}
