const N int = 26

func minimizeStringValue(s string) string {
    cnt := [N]int{}
    for i := range s {
        if s[i] == '?' {
            continue
        }
        cnt[int(s[i] - 'a')]++
    }

    h := &hp{}
    for i, x := range cnt {
        heap.Push(h, [2]int{i, x})
    }

    ans := []byte(s)
    tmp := []byte{}
    for i := range ans {
        if ans[i] == '?' {
            v := heap.Pop(h).([2]int)
            tmp = append(tmp, byte(v[0] + 'a'))
            v[1]++
            heap.Push(h, v)
        }
    }
    sort.Slice(tmp, func(i, j int) bool {
        return tmp[i] < tmp[j]
    })
    for i := range ans {
        if ans[i] == '?' {
            ans[i] = tmp[0]
            tmp = tmp[1:]
        }
    }

    return string(ans)
}

type hp [][2]int
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i][1] < h[j][1] || h[i][1] == h[j][1] && h[i][0] < h[j][0]}
func (h hp) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *hp) Push(x any) {
    *h = append(*h, x.([2]int))
}
func (h *hp) Pop() (x any) {
    old := *h 
    n := len(old)
    x = old[n - 1]
    *h = old[:n - 1]
    return
}