func maxDistinct(s string) int {
    mask := 0
    for _, c := range s {
        x := int(c - 'a')
        mask |= 1 << x
    }

    return bits.OnesCount(uint(mask))
}