func getSmallestString(s string, k int) string {
    t := []byte(s)
    for i := 0; i < len(t) && k > 0; i++ {
        t[i] = find(t[i], &k)
    }

    return string(t)
}

func find(c byte, k *int) byte {
    x := int(c - 'a')
    if *k >= min(x, 26 - x) {
        *k -= min(x, 26 - x)
        return 'a'
    }

    defer func() {
        *k = 0
    }()
    return byte('a' + x - *k)
}