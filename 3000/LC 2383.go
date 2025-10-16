func minNumberOfHours(a int, b int, energy []int, experience []int) (ans int) {
    n := len(energy)
    for i := range n {
        x, y := energy[i], experience[i]
        if a <= x {
            ans += x - a + 1
            a = x + 1
        } 
        if b <= y {
            ans += y - b + 1
            b = y + 1
        }

        a -= x 
        b += y
    }

    return 
}