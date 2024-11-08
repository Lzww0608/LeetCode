type LogSystem struct {
    id2Time map[string]int
    gran map[string]int
}


func Constructor() LogSystem {
    gran := map[string]int{
        "Year": 4,
        "Month": 7,
        "Day": 10,
        "Hour": 13,
        "Minute": 16,
        "Second": 19,
    }
    return LogSystem{map[string]int{}, gran}
}


func (c *LogSystem) Put(id int, timestamp string)  {
    c.id2Time[timestamp] = id
}


func (c *LogSystem) Retrieve(start string, end string, granularity string) (ans []int) {
    start = start[:c.gran[granularity]]
    end = end[:c.gran[granularity]]
    for t, x := range c.id2Time {
        t = t[:c.gran[granularity]]
        if t >= start && t <= end {
            ans = append(ans, x)
        }
    }

    return 
}


/**
 * Your LogSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(id,timestamp);
 * param_2 := obj.Retrieve(start,end,granularity);
 */