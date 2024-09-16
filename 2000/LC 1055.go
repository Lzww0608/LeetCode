import (
    "sort"
)
func shortestWay(s string, t string) (ans int) {
    cnt := make([][]int, 26)
    for i := range s {
        x := int(s[i] - 'a')
        cnt[x] = append(cnt[x], i)
    }

    pre := -1
    for i := range t {
        x := int(t[i] - 'a')
        if len(cnt[x]) == 0 {
            return -1
        }
        pos := sort.SearchInts(cnt[x], pre + 1)
        if pos == len(cnt[x]) {
            ans++
            pre = cnt[x][0]
        } else {
            pre = cnt[x][pos]
        }
    }

    return ans + 1
}