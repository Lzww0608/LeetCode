const mod = 769

type pair struct {
    k, v int
}

type MyHashMap struct {
    data []list.List
}


func Constructor() MyHashMap {
    return MyHashMap{make([]list.List, mod)}
}

func (m *MyHashMap) hash(key int) int {
    return key % mod;
}


func (m *MyHashMap) Put(key int, value int)  {
    h := m.hash(key)
    for e := m.data[h].Front(); e != nil; e = e.Next() {
        if et := e.Value.(pair); et.k == key {
            e.Value = pair{key, value}
            return
        }
    }
    m.data[h].PushBack(pair{key, value})
}


func (m *MyHashMap) Get(key int) int {
    h := m.hash(key)
    for e := m.data[h].Front(); e != nil; e = e.Next() {
        if et := e.Value.(pair); et.k == key {
            return et.v
        }
    }
    return -1
}


func (m *MyHashMap) Remove(key int)  {
    h := m.hash(key)
    for e := m.data[h].Front(); e != nil; e = e.Next() {
        if et := e.Value.(pair); et.k == key {
            m.data[h].Remove(e)
            return
        }
    }
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
