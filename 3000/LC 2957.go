func removeAlmostEqualCharacters(s string) (ans int) {
    n := len(s)
    for i := 1; i < n; i++ {
        if abs(int(s[i]) - int(s[i - 1])) <= 1 {
            ans++
            i++
        } 
    }

    return
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}