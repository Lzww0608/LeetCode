type Logger struct {
    m map[string]int
}


func Constructor() Logger {
    m := make(map[string]int)
    return Logger{m}
}


func (c *Logger) ShouldPrintMessage(timestamp int, message string) bool {
    if v, ok := c.m[message]; !ok || v + 10 <= timestamp {
        c.m[message] = timestamp
        return true
    }

    return false
}


/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */