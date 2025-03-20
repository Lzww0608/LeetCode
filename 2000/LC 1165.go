const N int = 26
func calculateTime(keyboard string, word string) (ans int) {
    pos := make([]int, N)
    for i := range keyboard {
        x := int(keyboard[i] - 'a')
        pos[x] = i 
    }

    cur := 0
    for i := range word {
        x := pos[int(word[i] - 'a')]
        ans += abs(cur - x)
        cur = x
    }

    return
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}