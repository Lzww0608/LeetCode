type FrequencyTracker struct {
    m, fre map[int]int
}


func Constructor() FrequencyTracker {
    m := map[int]int{}
    fre := map[int]int{}
    return FrequencyTracker{m, fre}
}


func (this *FrequencyTracker) Add(number int)  {
    this.fre[this.m[number]]--
    this.m[number]++
    this.fre[this.m[number]]++
}


func (this *FrequencyTracker) DeleteOne(number int)  {
    if this.m[number] > 0 {
        this.fre[this.m[number]]--
        this.m[number]--
        this.fre[this.m[number]]++
    }
}


func (this *FrequencyTracker) HasFrequency(frequency int) bool {
    if this.fre[frequency] > 0 {
        return true
    }
    return false
}


/**
 * Your FrequencyTracker object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * obj.DeleteOne(number);
 * param_3 := obj.HasFrequency(frequency);
 */