const N int = 26
func numberOfSpecialChars(word string) (ans int) {
    cnt := make([]int, 26)
    for i := range word {
        if word[i] >= 'a' && word[i] <= 'z' {
            x := int(word[i] - 'a')
            if cnt[x] == -1 {
                ans--
                cnt[x] = -2
            } else if cnt[x] != -2 {
                cnt[x]++
            }
        } else {
            x := int(word[i] - 'A')
            if cnt[x] > 0 {
                ans++
                cnt[x] = -1
            } else if cnt[x] == 0 {
                cnt[x] = -2
            }
        }
    }

    return
}