func generateTag(caption string) string {
    s := strings.Fields(caption)
    var ans strings.Builder
    ans.WriteByte('#')

    for i, t := range s {
        for j := range t {
            c := t[j]
            if j == 0 && i != 0 && unicode.IsLower(rune(c)) {
                c -= 32
            } else if (j > 0 || i == 0) && unicode.IsUpper(rune(c)) {
                c += 32
            }

            ans.WriteByte(c)
            if ans.Len() == 100 {
                return ans.String()
            }
        }
    }

    return ans.String()
}