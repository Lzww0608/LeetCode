/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func findRoot(tree []*Node) *Node {
    xor := 0
    for _, v := range tree {
        xor ^= v.Val
        for _, t := range v.Children {
            xor ^= t.Val
        }
    }

    for _, v := range tree {
        if xor == v.Val {
            return v
        }
    }

    return nil
}