type Robot struct {
    k int
    dirs [][]int
    m []string
    i, j int
    w, h int
}


func Constructor(width int, height int) Robot {
    dirs := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
    k := 0
    m := []string{"East", "North", "West", "South"}
    return Robot{k, dirs, m, 0, 0, width, height}
}


func (c *Robot) Step(num int)  {
    perimeter := 2*(c.w+c.h) - 4
    if perimeter == 0 {
        return
    }
    num %= perimeter

    if num == 0 && (c.i == 0 && c.j == 0) {
        c.k = 3 
    }

    for num > 0 {
        num--
        c.i += c.dirs[c.k][0]
        c.j += c.dirs[c.k][1]
        if c.i < 0 || c.i >= c.w || c.j < 0 || c.j >= c.h {
            c.i -= c.dirs[c.k][0]
            c.j -= c.dirs[c.k][1]
            c.k = (c.k + 1) % 4
            c.i += c.dirs[c.k][0]
            c.j += c.dirs[c.k][1]
        }
    }
}


func (c *Robot) GetPos() []int {
    return []int{c.i, c.j}
}


func (c *Robot) GetDir() string {
    return c.m[c.k]
}


/**
 * Your Robot object will be instantiated and called as such:
 * obj := Constructor(width, height);
 * obj.Step(num);
 * param_2 := obj.GetPos();
 * param_3 := obj.GetDir();
 */