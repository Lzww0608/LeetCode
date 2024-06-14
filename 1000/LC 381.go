type RandomizedCollection struct {
    m map[int]map[int]struct{}
    a []int
}


func Constructor() RandomizedCollection {
    return RandomizedCollection{map[int]map[int]struct{}{}, []int{}}
}


func (this *RandomizedCollection) Insert(val int) bool {
    v, ok := this.m[val]
    if !ok {
        v = map[int]struct{}{}
        this.m[val] = v    
    }
    v[len(this.a)] = struct{}{}
    this.a = append(this.a, val)
    return !ok
}


func (this *RandomizedCollection) Remove(val int) bool {
    v, ok := this.m[val]
    if !ok {
        return false
    }
    i := 0
    for x := range v {
        i = x
        break
    }
    n := len(this.a)
    this.a[i] = this.a[n-1]
    delete(v, i)
    delete(this.m[this.a[i]], n - 1)
    if i < n - 1 {
        this.m[this.a[i]][i] = struct{}{}
    }
    if len(v) == 0 {
        delete(this.m, val)
    }
    this.a = this.a[:n-1]
    return true
}


func (this *RandomizedCollection) GetRandom() int {
    return this.a[rand.Intn(len(this.a))]
}


/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */