type ExamTracker struct {
    time []int 
    pre []int
}


func Constructor() ExamTracker {
    return ExamTracker{[]int{}, []int{0}}
}


func (c *ExamTracker) Record(time int, score int)  {
    c.time = append(c.time, time)
    c.pre = append(c.pre, c.pre[len(c.pre) - 1] + score)
}


func (c *ExamTracker) TotalScore(startTime int, endTime int) int64 {
    i := sort.SearchInts(c.time, startTime)
    j := sort.SearchInts(c.time, endTime)
    if c.time[j] > endTime {
        j--
    }

    return int64(c.pre[j + 1] - c.pre[i])
}


/**
 * Your ExamTracker object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Record(time,score);
 * param_2 := obj.TotalScore(startTime,endTime);
 */