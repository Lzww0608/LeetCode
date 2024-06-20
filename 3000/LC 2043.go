type Bank struct {
    mny []int64
    cnt int
}


func Constructor(balance []int64) Bank {
    n := len(balance)
    return Bank{balance, n}
}


func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
    if account1 == 0 || account2 == 0 || account1 > this.cnt || account2 > this.cnt || this.mny[account1-1] < money {
        return false
    }
    this.mny[account1-1] -= money
    this.mny[account2-1] += money
    return true
}


func (this *Bank) Deposit(account int, money int64) bool {
    if account == 0 || account > this.cnt {
        return false
    }
    this.mny[account-1] += money
    return true
}


func (this *Bank) Withdraw(account int, money int64) bool {
    if account == 0 || account > this.cnt || this.mny[account-1] < money {
        return false
    }
    this.mny[account-1] -= money
    return true
}


/**
 * Your Bank object will be instantiated and called as such:
 * obj := Constructor(balance);
 * param_1 := obj.Transfer(account1,account2,money);
 * param_2 := obj.Deposit(account,money);
 * param_3 := obj.Withdraw(account,money);
 */
