type ATM struct {
    sum []int
}

var money = [5]int{20, 50, 100, 200, 500}


func Constructor() ATM {
    sum := make([]int, 5)
    return ATM{sum}
}


func (c *ATM) Deposit(banknotesCount []int)  {
    for i, x := range banknotesCount {
        c.sum[i] += x
    }
}


func (c *ATM) Withdraw(amount int) []int {
    ans := make([]int, 5)

    for i := 4; i >= 0; i-- {
        a := min(c.sum[i], amount / money[i])
        c.sum[i] -= a 
        ans[i] += a
        amount -= a * money[i]
    }

    if amount != 0 {
        for i, x := range ans {
            c.sum[i] += x
        }
        return []int{-1}
    }

    return ans
}




/**
 * Your ATM object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Deposit(banknotesCount);
 * param_2 := obj.Withdraw(amount);
 */