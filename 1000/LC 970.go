func powerfulIntegers(x int, y int, bound int) (ans []int) {
    m := map[int]bool{}

    for i := 1; i <= bound; i *= x {
        for j := 1; i + j <= bound; j *= y {
            if !m[i + j] {
                ans = append(ans, i + j)
                m[i + j] = true
            }
            if y == 1 {
                break
            }
        }
        if x == 1 {
            break
        }
    } 

    return 
}

