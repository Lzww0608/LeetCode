func discountPrices(sentence string, discount int) string {
    str := strings.Split(sentence, " ")
    builder := strings.Builder{}
    for _, s := range str {
        if s[0] == '$' && len(s) > 1 {
            x, i := 0, 1
            neg := false
            if i < len(s) && s[i] == '-' {
                neg = true
                i++
            }

            for ; i < len(s); i++ {
                if s[i] >= '0' && s[i] <= '9' {
                    x = x * 10 + int(s[i] - '0')
                } else {
                    break
                }
            }
            if i == len(s) {
                if neg {
                    x *= -1
                }
                t := float64(x) * float64(100 - discount) / 100.0
                builder.WriteString("$" + fmt.Sprintf("%.2f", t) + " ")
                continue
            }
        } 
        builder.WriteString(s + " ")
    }
    ans := builder.String()

    return ans[:len(ans)-1]
}