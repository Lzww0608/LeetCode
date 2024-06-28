func hasGroupsSizeX(deck []int) bool {
    m := map[int]int{}

    for _, x := range deck {
        m[x]++
    }


    x := -1
    for _, v := range m {
        if x == -1 {
            x = v
        } else {
            x = gcd(x, v)
        }
    }

    return x >= 2
}

func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x % y
    }

    return x
}
