const N int = 26
func checkAlmostEquivalent(word1 string, word2 string) bool {
    cnt := [N]int{}
    for _, c := range word1 {
        x := int(c - 'a')
        cnt[x]++
    }
    for _, c := range word2 {
        x := int(c - 'a')
        cnt[x]--
    }

    for _, x := range cnt {
        if x < -3 || x > 3 {
            return false
        }
    }

    return true
}