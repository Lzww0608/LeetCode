func vowelStrings(words []string, left int, right int) (ans int) {
    mask := 0x104111
    for _, s := range words[left:right+1] {
        x := int(s[0] - 'a')
        y := int(s[len(s)-1] - 'a')
        if ((1 << x) & mask) != 0 && ((1 << y) & mask) != 0 {
            ans++
        }
    }

    return
}