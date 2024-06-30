func minDeletionSize(strs []string) (ans int) {
    for i := 0; i < len(strs[0]); i++ {
        c := byte('a' - 1)
        for _, s := range strs {
            if s[i] < c {
                ans++
                break
            }
            c = s[i]
        }
    }

    return 
}
