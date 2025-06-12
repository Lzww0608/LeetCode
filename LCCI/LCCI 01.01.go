func isUnique(astr string) bool {
    mask := 0 
    for _, c := range astr {
        x := int(c - 'a')
        if mask & (1 << x) != 0 {
            return false
        }
        mask |= 1 << x
    } 
    
    return true
}