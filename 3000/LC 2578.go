const N int = 10
func splitNum(num int) int {
    s := []byte(strconv.Itoa(num))
    sort.Slice(s, func(i, j int) bool {
        return s[i] < s[j]
    })

    a, b := 0, 0
    for i, c := range s {
        x := int(c - '0')
        if i & 1 == 0 {
            a = a * 10 + x
        } else {
            b = b * 10 + x
        }
    }

    return a + b
}