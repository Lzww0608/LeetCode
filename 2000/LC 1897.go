const N int = 26
func makeEqual(words []string) bool {
    n := len(words)
    cnt := [N]int{}
    for _, word := range words {
        for _, c := range word {
            cnt[int(c - 'a')]++
        }
    }

    for _, x := range cnt {
        if x % n != 0 {
            return false
        }
    }

    return true
}