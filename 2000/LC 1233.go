
import "strings"
func removeSubfolders(folder []string) (ans []string) {
    sort.Strings(folder)
    last := "#"
    for _, v := range folder {
        if !strings.HasPrefix(v, last) || v[len(last)] != '/' {
            ans = append(ans, v)
            last = v
        }
    }

    return
}