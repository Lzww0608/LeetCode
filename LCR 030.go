//惰性删除

type RandomizedSet struct {
    m map[int]int
    arr [200005]int
}

var idx int = 0

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
    idx = 0
    m := map[int]int{}
    arr := [200005]int{}
    return RandomizedSet{m, arr}
}


/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
    if _, ok := this.m[val]; ok {
        return false
    }
    this.m[val] = idx
    this.arr[idx] = val
    idx++
    return true
}


/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
    if _, ok := this.m[val]; !ok {
        return false
    }
    idx--
    t, x := this.m[val], this.arr[idx]
    this.arr[t] = x
    this.m[x] = t
    delete(this.m, val)
    return true
}


/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
    t := rand.Intn(idx)
    return this.arr[t]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */