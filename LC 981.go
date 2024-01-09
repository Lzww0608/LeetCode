type TimeMap struct {
    m map[string][]Node
}
type Node struct{
    t int
    s string
}

func Constructor() TimeMap {
    return TimeMap{make(map[string][]Node)}
}


func (this *TimeMap) Set(key string, value string, timestamp int)  {
    if _, ok := this.m[key]; !ok {
        this.m[key] = make([]Node, 0)
    }
    this.m[key] = append(this.m[key], Node{timestamp, value})
}


func (this *TimeMap) Get(key string, timestamp int) string {
    if _, ok := this.m[key]; !ok {
        return ""
    }
    nodes := this.m[key]
    left, right := 0, len(nodes) - 1

    // 执行二分搜索
    res := -1
    for left <= right {
        mid := left + (right-left)/2
        if nodes[mid].t <= timestamp {
            res = mid
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    // 如果找到合适的节点，返回对应的值
    if res != -1 {
        return nodes[res].s
    }
    return ""
}



/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */