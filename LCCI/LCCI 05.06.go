func convertInteger(A int, B int) (ans int) {
    for i := 0; i < 32; i++ {
        x := (A >> i) & 1
        y := (B >> i) & 1
        if x != 0 && y == 0 || x == 0 && y != 0 {
            ans++
        }

    }

    return
}