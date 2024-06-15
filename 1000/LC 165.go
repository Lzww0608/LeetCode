func compareVersion(version1 string, version2 string) int {
    s := strings.Split(version1, ".")
    t := strings.Split(version2, ".")
    i, j := 0, 0
    n, m := len(s), len(t)
    for i < n && j < m {
        a, _ := strconv.Atoi(s[i])
        b, _ := strconv.Atoi(t[j])
        if a < b {
            return -1
        } else if a > b {
            return 1
        }
        i++
        j++
    }
    for i < n {
        a, _ := strconv.Atoi(s[i])
        if a > 0 {
            return 1
        }
        i++
    }
    for j < m {
        a, _ := strconv.Atoi(t[j])
        if a > 0 {
            return -1
        }
        j++
    }

    return 0
}
