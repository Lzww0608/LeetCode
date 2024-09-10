const MOD int = 1_000_000_007
type Fancy struct {
    v []int
    a, b int
}

func quickPow(a, r int) int {
    res := 1
    for r > 0 {
        if r & 1 == 1 {
            res = (res * a) % MOD
        }
        r >>= 1
        a = a * a % MOD
    }

    return res
}

func inv(x int) int {
    return quickPow(x, MOD - 2)
}


func Constructor() Fancy {
    return Fancy{[]int{}, 1, 0}
}


func (c *Fancy) Append(val int)  {
    c.v = append(c.v, (val - c.b + MOD) % MOD * inv(c.a) % MOD)
}


func (c *Fancy) AddAll(inc int)  {
    c.b = (c.b + inc) % MOD
}


func (c *Fancy) MultAll(m int)  {
    c.a = c.a * m % MOD
    c.b = c.b * m % MOD
}


func (c *Fancy) GetIndex(idx int) int {
    if idx >= len(c.v) {
        return -1
    }

    ans := (c.a * c.v[idx] % MOD + c.b)% MOD

    return ans
}


/**
 * Your Fancy object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Append(val);
 * obj.AddAll(inc);
 * obj.MultAll(m);
 * param_4 := obj.GetIndex(idx);
 */