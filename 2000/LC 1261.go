/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type FindElements struct {
    root *TreeNode
}


func Constructor(root *TreeNode) FindElements {
    return FindElements{root}
}


func (this *FindElements) Find(target int) bool {
    cur := this.root
    target++
    for i := bits.Len(uint(target)) - 2; i >= 0; i-- {
        b := target >> i & 1
        if b == 1 {
            cur = cur.Right
        } else {
            cur = cur.Left
        }
        
        if cur == nil {
            return false
        }
    }
    return true
}


/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */
