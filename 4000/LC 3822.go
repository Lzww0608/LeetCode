type OrderManagementSystem struct {
    a map[int]pair
    b map[pair][]int
}

type pair struct {
    orderType string 
    price int
}


func Constructor() OrderManagementSystem {
    a := map[int]pair{}
    b := map[pair][]int{}
    return OrderManagementSystem{a, b}
}


func (c *OrderManagementSystem) AddOrder(orderId int, orderType string, price int)  {
    p := pair{orderType, price}
    c.a[orderId] = p
    c.b[p] = append(c.b[p], orderId)
}


func (c *OrderManagementSystem) ModifyOrder(orderId int, newPrice int)  {
    p := c.a[orderId]
    p.price = newPrice
    c.a[orderId] = p 
    c.b[p] = append(c.b[p], orderId)
}


func (c *OrderManagementSystem) CancelOrder(orderId int)  {
    delete(c.a, orderId)
}


func (c *OrderManagementSystem) GetOrdersAtPrice(orderType string, price int) (ans []int) {
    p := pair{orderType, price}
    vis := make(map[int]bool)
    for _, x := range c.b[p] {
        if c.a[x] == p && !vis[x] {
            ans = append(ans, x)
            vis[x] = true
        }
    }

    return
}


/**
 * Your OrderManagementSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddOrder(orderId,orderType,price);
 * obj.ModifyOrder(orderId,newPrice);
 * obj.CancelOrder(orderId);
 * param_4 := obj.GetOrdersAtPrice(orderType,price);
 */