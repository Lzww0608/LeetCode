
import "github.com/emirpasic/gods/trees/redblacktree"
type NumberContainers struct {
    num2Pos map[int]*redblacktree.Tree
    pos2Num map[int]int
}


func Constructor() NumberContainers {
    return NumberContainers{
		num2Pos: make(map[int]*redblacktree.Tree),
		pos2Num: make(map[int]int),
	}
}


func (c *NumberContainers) Change(index int, number int)  {
    if v, ok := c.pos2Num[index]; ok {
        c.num2Pos[v].Remove(index)
    } 
    c.pos2Num[index] = number
    if _, ok := c.num2Pos[number]; !ok {
        c.num2Pos[number] = redblacktree.NewWithIntComparator()
    }
    c.num2Pos[number].Put(index, nil)
}


func (c *NumberContainers) Find(number int) int {
    tree, exists := c.num2Pos[number]
	if !exists || tree.Size() == 0 {
		return -1
	}

	return tree.Left().Key.(int)
}


/**
 * Your NumberContainers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Change(index,number);
 * param_2 := obj.Find(number);
 */