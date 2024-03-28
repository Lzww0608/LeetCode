func reverseBits(num uint32) (x uint32) {
    for i := 0; i < 32; i++ {
        x = x << 1 + num & 1
        num >>= 1
    }
    return 
}