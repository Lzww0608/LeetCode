func crackPassword(password []int) string {

    cmp := func(x, y int) bool {
        a := strconv.Itoa(x)
        b := strconv.Itoa(y)
        return  a + b < b + a
    }

    sort.Slice(password, func(i, j int) bool {
        return cmp(password[i], password[j])
    })

    builder := strings.Builder{}
    for _, x := range password {
        builder.WriteString(strconv.Itoa(x))
    }
    return builder.String()
}