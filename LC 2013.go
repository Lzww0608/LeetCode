type DetectSquares struct {
    m map[[2]int]int
    cnt [1005][1005]int
}


func Constructor() DetectSquares {
    return DetectSquares{map[[2]int]int{}, [1005][1005]int{}}
}


func (this *DetectSquares) Add(point []int)  {
    x, y := point[0], point[1]
    this.m[[2]int{x, y}]++
    this.cnt[y][x]++
}


func (this *DetectSquares) Count(point []int) (ans int) {
    x, y := point[0], point[1]
    for i := 0; i <= 1000; i++ {
        if x != i && this.cnt[y][i] > 0 {
            d := x - i
            ans += this.m[[2]int{i, y}] * this.m[[2]int{i, y + d}] * this.m[[2]int{x, y + d}]
            ans += this.m[[2]int{i, y}] * this.m[[2]int{i, y - d}] * this.m[[2]int{x, y - d}]
        }
    }

    return
}



/**
 * Your DetectSquares object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(point);
 * param_2 := obj.Count(point);
 */