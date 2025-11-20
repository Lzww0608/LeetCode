const N int = 26
func numberOfSpecialChars(word string) (ans int) {
    mask := [N]int{}
    for _, c := range word {
        if unicode.IsLower(c) {
            x := int(c - 'a')
            mask[x] |= 1
        } else if unicode.IsUpper(c) {
            x := int(c - 'A')
            mask[x] |= 2
        }
    }

    for _, x := range mask {
        if x == 3 {
            ans++
        }
    }

    return
}