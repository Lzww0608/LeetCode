type Solution struct {
    bound int
    m map[int]int
}


func Constructor(n int, blacklist []int) Solution {
    exist := map[int]bool{}
    for _, x := range blacklist {
        exist[x] = true
    }
    m := map[int]int{}
    k, i := n - len(blacklist), n - len(blacklist)
    for _, x := range blacklist {
        if x < k {
            for exist[i] {
                i++
            }
            m[x] = i 
            i++
        }
        
    }
    return Solution{k, m}
}


func (this *Solution) Pick() int {
    i := rand.Intn(this.bound)
    if v, ok := this.m[i]; ok {
        return v
    }
    return i
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n, blacklist);
 * param_1 := obj.Pick();
 */