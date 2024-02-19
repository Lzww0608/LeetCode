type FreqStack struct {
    maxFrequency int
    frequencyMap map[int]int 
    stacksOfFrequency [20000][]int
}


func Constructor() FreqStack {
    return FreqStack{0, map[int]int{}, [20000][]int{}}
}


func (this *FreqStack) Push(val int)  {
    this.frequencyMap[val]++
    if this.frequencyMap[val] > this.maxFrequency {
        this.maxFrequency = this.frequencyMap[val]
    } 
    this.stacksOfFrequency[this.frequencyMap[val]] = append(this.stacksOfFrequency[this.frequencyMap[val]], val)

}


func (this *FreqStack) Pop() int {
    for len(this.stacksOfFrequency[this.maxFrequency]) == 0 {
        this.maxFrequency--
    }
    res := this.stacksOfFrequency[this.maxFrequency][len(this.stacksOfFrequency[this.maxFrequency])-1]
    this.stacksOfFrequency[this.maxFrequency] = this.stacksOfFrequency[this.maxFrequency][:len(this.stacksOfFrequency[this.maxFrequency])-1]
    this.frequencyMap[res]--
    return res
}


/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */