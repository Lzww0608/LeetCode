func minFlips(a int, b int, c int) (ans int) {
    for i := 1; i < (1 << 31); i <<= 1 {
        x, y, z := a & i, b & i, c & i 
        if (x | y == 0) && (z != 0) {
            ans++
        } else if (x | y != 0) && (z == 0) {
            if x != 0 {
                ans++
            }
            if y != 0 {
                ans++
            }
        }
    }

    return
}