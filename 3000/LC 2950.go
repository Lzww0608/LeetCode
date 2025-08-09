func countDivisibleSubstrings(s string) (ans int) {
    m := map[byte]int {
        'a': 1,
        'b': 1,
        'c': 2,
        'd': 2,
        'e': 2,
        'f': 3,
        'g': 3,
        'h': 3,
        'i': 4,
        'j': 4,
        'k': 4,
        'l': 5,
        'm': 5,
        'n': 5,
        'o': 6,
        'p': 6,
        'q': 6,
        'r': 7,
        's': 7,
        't': 7,
        'u': 8,
        'v': 8,
        'w': 8,
        'x': 9,
        'y': 9,
        'z': 9,
    }

    for avg := 1; avg <= 9; avg++ {
        cnt := make(map[int]int)
        cnt[0] = 1
        sum := 0
        for i := range s {
            sum += m[s[i]] - avg
            ans += cnt[sum]
            cnt[sum]++
        }
    }

    return 
}