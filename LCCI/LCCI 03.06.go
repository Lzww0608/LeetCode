type AnimalShelf struct {
    stCat, stDog [][]int
    idx int
}


func Constructor() AnimalShelf {
    return AnimalShelf{[][]int{}, [][]int{}, 0}
}


func (c *AnimalShelf) Enqueue(animal []int)  {
    n, t := animal[0], animal[1]
    if t == 0 {
        c.stCat = append(c.stCat, []int{n, c.idx})
    } else {
        c.stDog = append(c.stDog, []int{n, c.idx})
    }
    c.idx++
}


func (c *AnimalShelf) DequeueAny() (ans []int) {
    if len(c.stCat) == 0 && len(c.stDog) == 0 {
        return []int{-1, -1}
    } else if len(c.stDog) == 0 || (len(c.stCat) > 0 && c.stCat[0][1] < c.stDog[0][1]) {
        ans = []int{c.stCat[0][0], 0}
        c.stCat = c.stCat[1:]
        return
    } 
    ans = []int{c.stDog[0][0], 1}
    c.stDog = c.stDog[1:]
    return
}


func (c *AnimalShelf) DequeueDog() (ans []int) {
    if len(c.stDog) == 0 {
        return []int{-1, -1}
    }
    ans = []int{c.stDog[0][0], 1}
    c.stDog = c.stDog[1:]
    return
}


func (c *AnimalShelf) DequeueCat() (ans []int) {
    if len(c.stCat) == 0 {
        return []int{-1, -1}
    }
    ans = []int{c.stCat[0][0], 0}
    c.stCat = c.stCat[1:]
    return
}


/**
 * Your AnimalShelf object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Enqueue(animal);
 * param_2 := obj.DequeueAny();
 * param_3 := obj.DequeueDog();
 * param_4 := obj.DequeueCat();
 */