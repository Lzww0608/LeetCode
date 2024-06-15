func minDeletions(s string) (ans int) {
    freq := make([]int, 26)
    for _, c := range s {
        freq[int(c - 'a')]++
    }
    sort.Slice(freq, func(i, j int) bool {
        return freq[i] > freq[j]
    })
    m := map[int]int{}
    for _, x := range freq {
        if x > 0 {
            m[x]++
        }
    }

    prev := freq[0]
    for i := 1; i < 26 && freq[i] > 0; i++ {
        if prev <= freq[i] {
            prev = max(prev - 1, 0)
            ans += freq[i] - prev
        } else {
            prev = freq[i]
        }
        
    }

    return
}