
type Cashier struct {
    n, discount, id int
    m map[int]int
}


func Constructor(n int, discount int, products []int, prices []int) Cashier {
    m := map[int]int{}
    for i := range products {
        m[products[i]] = prices[i]
    }
    return Cashier{n, discount, 0, m}
}


func (c *Cashier) GetBill(product []int, amount []int) (ans float64) {
    for i := range product {
        ans += float64(c.m[product[i]] * amount[i])
    }

    c.id++
    if c.id % c.n == 0 {
        ans *= float64(100 - c.discount) / 100.0
    }

    return
}


/**
 * Your Cashier object will be instantiated and called as such:
 * obj := Constructor(n, discount, products, prices);
 * param_1 := obj.GetBill(product,amount);
 */