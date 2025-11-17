func canBeEqual(a string, b string) bool {
    s1 := []byte(a)
    s2 := []byte(b)
    if s1[0] > s1[2] {
        s1[0], s1[2] = s1[2], s1[0]
    }
    if s1[1] > s1[3] {
        s1[1], s1[3] = s1[3], s1[1]
    }

    if s2[0] > s2[2] {
        s2[0], s2[2] = s2[2], s2[0]
    }
    if s2[1] > s2[3] {
        s2[1], s2[3] = s2[3], s2[1]
    }

    return string(s1) == string(s2)
}