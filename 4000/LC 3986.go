func secondsBetweenTimes(startTime string, endTime string) (ans int) {
    l := strings.Split(startTime, ":")
    r := strings.Split(endTime, ":")

    a := [3]int{}
    b := [3]int{}
    for i := range a {
        a[i], _ = strconv.Atoi(l[i])
        b[i], _ = strconv.Atoi(r[i])
    }
    
    if a[0] == b[0] {
        ans = 60 * (b[1] - a[1]) + b[2] - a[2]
    } else {
        ans += 3600 * (b[0] - a[0] - 1)
        ans += 60 * b[1] + b[2]
        ans += 60 * (59 - a[1]) + 60 - a[2]
    }

    return
}