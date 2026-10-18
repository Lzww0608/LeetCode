type BusBooking struct {
    middle, window, t int 
    a [][4]int 
    n int 
} 


func Constructor(n int) BusBooking {
    a := make([][4]int, n)
    return BusBooking {
        middle: 0,
        window: 0,
        t: 0,
        n: n,
        a: a,
    }
}


func (b *BusBooking) Toggle(s string)  {
    x := int(s[0] - 'A') 
    y, _ := strconv.Atoi(s[1:])
    y--
    if b.a[y][x] != 0 {
        b.t -= b.a[y][x]
        b.a[y][x] = 0 
        if x == 1 || x == 2 {
            b.middle--
        } else {
            b.window--
        }
        return
    }
    
    if x == 1 || x == 2 {
        b.t++
        b.middle++
        b.a[y][x] = 1
    } else {
        if x == 0 && b.a[y][x + 1] != 0 || x == 3 && b.a[y][x - 1] != 0 {
            b.t += 3
            b.a[y][x] = 3
        } else {
            b.t += 1
            b.a[y][x] = 1
        }
        b.window++
    }
    return
}


func (b *BusBooking) GetTotalTime() int {
    return b.t
}


func (b *BusBooking) GetMinTime() int {
    return b.middle + b.window
}


/**
 * Your BusBooking object will be instantiated and called as such:
 * obj := Constructor(n);
 * objToggle(seat);
 * param_2 := obj.GetTotalTime();
 * param_3 := obj.GetMinTime();
 */