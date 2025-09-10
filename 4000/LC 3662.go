
import "strings"
const N int = 26
func filterCharacters(s string, k int) string {
    cnt := [N]int{}
    for i := range s {
        x := int(s[i] - 'a')
        cnt[x]++
    }

    ans := strings.Builder{}
    for i := range s {
        x := int(s[i] - 'a')
        if cnt[x] < k {
            ans.WriteByte(s[i])
        }
    }

    return ans.String()
}