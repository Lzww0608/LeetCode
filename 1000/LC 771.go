func numJewelsInStones(jewels string, stones string) int {
    mask := 0
    for _, c := range jewels {
        mask |= 1 << (c & 63)
    }

    ans := 0
    for _, c := range stones {
        ans += mask >> (c & 63) & 1
    }

    return ans
}
