/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
    st []*TreeNode
}


func Constructor(root *TreeNode) BSTIterator {
    st := []*TreeNode{}
    for cur := root; cur != nil; cur = cur.Left {
        st = append(st, cur)
    }
    return BSTIterator{st}
}


func (c *BSTIterator) Next() (ans int) {
    cur := c.st[len(c.st)-1]
    ans = cur.Val
    c.st = c.st[:len(c.st)-1]

    for p := cur.Right; p != nil; p = p.Left {
        c.st = append(c.st, p)
    }

    return
}


func (c *BSTIterator) HasNext() bool {
    return len(c.st) > 0
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */