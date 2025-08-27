func findMinDifference(timePoints []string) int {
    if len(timePoints) > 24 * 60 {
        return 0
    }
    a := []int{}
    for _, time := range timePoints {
        s := strings.Split(time, ":")
        h := int(s[0][0] - '0') * 10 + int(s[0][1] - '0')
        m := int(s[1][0] - '0') * 10 + int(s[1][1] - '0')
        a = append(a, h * 60 + m)
    }
    ans := math.MaxInt32
    sort.Ints(a)
    for i := 1; i < len(a); i++ {
        ans = min(ans, a[i] - a[i-1])
    }

    return min(ans, a[0] + 24 * 60 - a[len(a)-1])
}
