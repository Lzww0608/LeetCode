type Operations struct {
    neg, pos [31]int
}


func Constructor() Operations {
    var neg, pos [31]int
    a, b := -1, 1
    for i := 0; i < 31; i++ {
        neg[i] = a
        a += a
        pos[i] = b
        b += b
    }
    return Operations{neg, pos}
}

func (c *Operations) negate(a int) (res int) {
    if a == 0 {
        return 0
    }

    if a > 0 {
        for i := 30; i >= 0; i-- {
            if a + c.neg[i] >= 0 {
                a += c.neg[i]
                res += c.neg[i]
            }
        }
    } else {
        for i := 30; i >= 0; i-- {
            if a + c.pos[i] <= 0 {
                a += c.pos[i]
                res += c.pos[i]
            }
        }
    }

    return 
}


func (this *Operations) Minus(a int, b int) int {
    return a + this.negate(b)
}


func (this *Operations) Multiply(a int, b int) int {
    if a == 0 || b == 0 {
        return 0
    }

    if a == 1 {
        return b
    } else if b == 1 {
        return a
    }

    if b < 0 {
        a, b = this.negate(a), this.negate(b)
    }

    result := 0
    for b > 0 {
        if b&1 == 1 {
            result += a
        }
        a += a
        b >>= 1
    }

    return result
}


func (this *Operations) Divide(a int, b int) int {
    if a == 0 || b == 1 {
        return a
    } else if a == b {
        return 1
    } else if b == -1 {
        return this.negate(a)
    }

    if a > 0 {
        if b < 0 {
            return this.negate(this.Divide(a, this.negate(b)))
        } else if b > a {
            return 0
        }

        res, sum := 1, b
        for sum + sum > 0 && sum + sum < a {
            res += res
            sum += sum
        }

        return res + this.Divide(this.Minus(a, sum), b)
    } else {
        if b > 0 {
            return this.negate(this.Divide(a, this.negate(b))) 
        } else if b < a {
            return 0
        }

        res, sum := 1, b
        for sum + sum < 0 && sum + sum > a {
            res += res
            sum += sum
        }

        return res + this.Divide(this.Minus(a, sum), b)
    }
    return 0
}


/**
 * Your Operations object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Minus(a,b);
 * param_2 := obj.Multiply(a,b);
 * param_3 := obj.Divide(a,b);
 */