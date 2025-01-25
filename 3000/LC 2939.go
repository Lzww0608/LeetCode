const MOD = 1_000_000_007
func maximumXorProduct(a int64, b int64, n int) int {
    a, b = max(a, b), min(a, b)

    mask := int64(1 << n) - 1
    ax := a &^ mask
    bx := b &^ mask
    a &= mask
    b &= mask

    l := a ^ b 
    one := mask ^ l 
    ax |= one
    bx |= one

    if l > 0 && ax == bx {
        var h int64 = 1 << (bits.Len(uint(l)) - 1)
        ax |= h 
        l ^= h 
    }
    bx |= l

    return int(ax % MOD * (bx % MOD) % MOD)
}