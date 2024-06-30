func hammingDistance(x int, y int) (ans int) {
    s := x ^ y
    for s != 0 {
        s &= s - 1
        ans++
    }

    return
}
