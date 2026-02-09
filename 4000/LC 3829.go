type RideSharingSystem struct {
    riders []int 
    drivers []int
    m map[int]bool
}


func Constructor() RideSharingSystem {
    return RideSharingSystem {
        riders: []int{},
        drivers: []int{},
        m: map[int]bool{},
    }
}


func (c *RideSharingSystem) AddRider(riderId int) {
    c.riders = append(c.riders, riderId)
    c.m[riderId] = true
}


func (c *RideSharingSystem) AddDriver(driverId int) {
    c.drivers = append(c.drivers, driverId)
}


func (c *RideSharingSystem) MatchDriverWithRider() []int {
    for len(c.riders) > 0 {
        if _, ok := c.m[c.riders[0]]; !ok {
            c.riders = c.riders[1:]
        } else {
            break
        }
    }

    if len(c.drivers) == 0 || len(c.riders) == 0 {
        return []int{-1, -1}
    }
    a := c.riders[0]
    c.riders = c.riders[1:]
    b := c.drivers[0]
    c.drivers = c.drivers[1:]
    return []int{b, a}
}


func (c *RideSharingSystem) CancelRider(riderId int)  {
    delete(c.m, riderId)
}


/**
 * Your RideSharingSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddRider(riderId);
 * obj.AddDriver(driverId);
 * param_3 := obj.MatchDriverWithRider();
 * obj.CancelRider(riderId);
 */