func smallestNumber(pattern string) string {
    st := []int{}
    pattern += "I"
    builder := strings.Builder{}
    for i, c := range pattern {
        if c == 'I' {
            builder.WriteString(strconv.Itoa(i + 1))
            for len(st) > 0 {
                builder.WriteString(strconv.Itoa(st[len(st)-1]))
                st = st[:len(st)-1]
            } 
        } else {
            st = append(st, i + 1)
        }
    }
    


    return builder.String()
}