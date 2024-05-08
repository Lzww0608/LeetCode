func distinctSubseqII(s string) (res int) {
    mod := int(1e9 + 7)
    f := make([]int, 26)
    
    for i := range s {
        x := int(s[i] - 'a')
        sum := 0
        for _, v := range f {
            sum += v
        }
        f[x] = 1 + sum % mod
    }

    for _, x := range f {
        res = (res + x) % mod
    }

    return 
}