func hammingWeight(num uint32) (ans int) {
    for num != 0 {
        ans += int(num & 1)
        num >>= 1
    }

    return
}
