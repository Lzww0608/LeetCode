func stoneGameIX(stones []int) bool {
    a := [3]int{}
    for _, x := range stones {
        a[x % 3]++
    }

    if a[0] & 1 == 0 {
        return a[1] > 0 && a[2] > 0
    }

    return a[1] - 2 > a[2] || a[2] - 2 > a[1]
}



