func mapWordWeights(words []string, weights []int) string {
    ans := []byte{}
    for _, s := range words {
        x := 0
        for _, c := range s {
            x += weights[int(c - 'a')]
        }

        ans = append(ans, byte('a' + 25 - x % 26))
    }

    return string(ans)
}