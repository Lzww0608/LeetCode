func sockCollocation(sockets []int) []int {
    mask := 0
    for _, x := range sockets {
        mask ^= x
    }

    k := mask & (-mask)

    a, b := 0, 0
    for _, x := range sockets {
        if x & k == 0 {
            a ^= x
        } else {
            b ^= x
        }
    } 

    return []int{a, b}
}
