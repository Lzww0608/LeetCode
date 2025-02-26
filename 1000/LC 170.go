
import "math"
type TwoSum struct {
    cnt map[int]int
    mx, mn int
}


func Constructor() TwoSum {
    cnt := make(map[int]int)
    mn, mx := math.MaxInt32, math.MinInt32
    return TwoSum{
        cnt, mn, mx,
    }
}


func (c *TwoSum) Add(number int)  {
    c.cnt[number]++
    c.mn = min(c.mn, number)
    c.mx = max(c.mx, number)
}


func (c *TwoSum) Find(value int) bool {
    if value > c.mx * 2 || value < c.mn * 2{
        return false
    }

    for k, x := range c.cnt {
        if k * 2 == value {
            if x > 1 {
                return true
            }
        } else if c.cnt[value-k] > 0 {
            return true
        }
    }

    return false
}


/**
 * Your TwoSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * param_2 := obj.Find(value);
 */