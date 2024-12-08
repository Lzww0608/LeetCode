type SnakeGame struct {
    score int
    food [][]int
    m, n int
    q [][2]int
}


func Constructor(width int, height int, food [][]int) SnakeGame {
    return SnakeGame{0, food, height, width, [][2]int{{0, 0}}}
}


func (c *SnakeGame) Move(direction string) int {
    cur := c.q[len(c.q)-1]
    i, j := cur[0], cur[1]
    switch direction {
        case "U":
            i--
        case "D":
            i++
        case "L":
            j--
        case "R":
            j++
    }

    if i < 0 || i >= c.m || j < 0 || j >= c.n {
        return -1
    }

    if len(c.food) > 0 && c.food[0][0] == i && c.food[0][1] == j {
        c.score++
        c.food = c.food[1:]
    } else {
        c.q = c.q[1:]
    }

    for _, v := range c.q {
        if i == v[0] && j == v[1] {
            return -1
        }
    }

    c.q = append(c.q, [2]int{i, j})

    return c.score
}


/**
 * Your SnakeGame object will be instantiated and called as such:
 * obj := Constructor(width, height, food);
 * param_1 := obj.Move(direction);
 */