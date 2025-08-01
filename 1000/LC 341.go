/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedIterator struct {
    st [][]*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
    st := [][]*NestedInteger{nestedList}
    return &NestedIterator{st}
}

func (c *NestedIterator) Next() (ans int) {
    cur := c.st[len(c.st) - 1]
    ans = cur[0].GetInteger()
    c.st[len(c.st) - 1] = cur[1:]
    return
}

func (c *NestedIterator) HasNext() bool {
    for len(c.st) > 0 {
        cur := c.st[len(c.st) - 1]
        if len(cur) == 0 {
            c.st = c.st[:len(c.st)-1]
            continue
        }
        if cur[0].IsInteger() {
            return true
        }
        c.st[len(c.st) - 1] = cur[1:]
        c.st = append(c.st, cur[0].GetList())
    } 
    return false
}