/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
    head *ListNode
}


func Constructor(head *ListNode) Solution {
    return Solution {head}
}


func (c *Solution) GetRandom() (ans int) {
    k := 1
    for p := c.head; p != nil; p = p.Next {
        if rand.Intn(k) == 0 {
            ans = p.Val
        }
        k++
    }

    return 
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */