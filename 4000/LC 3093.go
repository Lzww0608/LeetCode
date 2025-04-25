
import "math"
func stringIndices(wordsContainer []string, wordsQuery []string) []int {
    type Node struct {
        son [26]*Node
        minL, i int
    }

    root := &Node{minL: math.MaxInt32}
    for j, s := range wordsContainer {
        n := len(s)
        cur := root
        if cur.minL > n {
            cur.minL = n 
            cur.i = j
        }
        for i := n - 1; i >= 0; i-- {
            x := int(s[i] - 'a')
            if cur.son[x] == nil {
                cur.son[x] = &Node{minL: math.MaxInt32, i: -1}
            }
            cur = cur.son[x]
            if cur.minL > n {
                cur.minL = n 
                cur.i = j
            }
        }
    }    
    //fmt.Println(root.i)
    ans := make([]int, len(wordsQuery))
    for j, s := range wordsQuery {
        cur := root
        ans[j] = cur.i
        for i := len(s) - 1; i >= 0; i-- {
            x := int(s[i] - 'a')
            cur = cur.son[x]
            if cur == nil {
                break
            }
            ans[j] = cur.i
        }
    }

    return ans
}