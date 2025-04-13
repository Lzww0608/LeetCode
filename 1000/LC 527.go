type Node struct {
    cnt int 
    children [26]*Node
}

func wordsAbbreviation(words []string) []string {
    n := len(words)
    ans := make([]string, n)
    type pair struct {
        s string 
        idx int
    }
    m := make(map[string][]pair)
    for i, s := range words {
        t := calc(s, 0)
        m[t] = append(m[t], pair{s, i})
    }

    //fmt.Println(m)
    for _, v := range m {
        root := &Node{}
        for _, t := range v {
            cur := root
            for i := 0; i < len(t.s); i++ {
                x := int(t.s[i] - 'a')
                if cur.children[x] == nil {
                    cur.children[x] = &Node{}
                }
                cur = cur.children[x]
                cur.cnt++
            }
        }

        for _, t := range v {
            cur := root
            var i int
            for i = 0; i < len(t.s); i++ {
                x := int(t.s[i] - 'a')
                cur = cur.children[x]
                if cur.cnt == 1 {
                    break
                }
            }
            ans[t.idx] = calc(t.s, i)
        }
    }

    return ans
}


func calc(s string, pos int) (ans string) {
    n := len(s)
    if pos + 3 >= n {
        return s 
    } 

    return s[:pos+1] + strconv.Itoa(n - pos - 2) + s[n-1:]
}