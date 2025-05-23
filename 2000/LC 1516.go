/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func moveSubTree(root *Node, p *Node, q *Node) *Node {
   var pp, qq *Node 
   isParent := false
   underQ := false

   var dfs func(cur, fa *Node, underP bool)
   dfs = func(cur, fa *Node, underP bool) {
        if cur == nil {
            return
        } 

        if cur == p {
            pp = fa
            underP = true
            if fa == q {
                underQ = true
                return
            }
        } else if cur == q {
            if underP {
                isParent = true
            }
            qq = fa 
        }
        for _, v := range cur.Children {
            dfs(v, cur, underP)
        }
    }

    dummy := &Node{-1, nil};
    dummy.Children = append(dummy.Children, root)
    dfs(root, dummy, false)

    if !underQ {
        if !isParent {
            n := len(pp.Children)
            for i := 0; i < n; i++ {
                if pp.Children[i] == p {
                    for j := i + 1; j < n; j++ {
                        pp.Children[j-1] = pp.Children[j]
                    }
                    break
                }
            }
            pp.Children = pp.Children[:n-1]
            q.Children = append(q.Children, p)
        } else {
            n := len(pp.Children)
            for i := 0; i < n; i++ {
                if pp.Children[i] == p {
                    pp.Children[i] = q
                    break
                }
            }
            n = len(qq.Children)
            for i := 0; i < n; i++ {
                if qq.Children[i] == q {
                    for j := i + 1; j < n; j++ {
                        qq.Children[j-1] = qq.Children[j]
                    }
                    break
                }
            }
            qq.Children = qq.Children[:n-1]
            q.Children = append(q.Children, p)
        }
    }

    return dummy.Children[0]
}