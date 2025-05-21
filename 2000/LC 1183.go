func maximumNumberOfOnes(width int, height int, sideLength int, maxOnes int) (ans int) {
    a := []int{}
    for i := 0; i < sideLength; i++ {
        for j := 0; j < sideLength; j++ {
            w := (width - i - 1) / sideLength + 1;
            d := (height - j - 1) / sideLength + 1;
            a= append(a, w * d)
        }
    }

    sort.Ints(a)
    for i := 0; i < maxOnes; i++ {
        ans += a[len(a)-i-1]
    }

    return
}