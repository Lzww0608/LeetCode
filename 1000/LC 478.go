type Solution struct {
    radius, x_center, y_center float64
}


func Constructor(radius float64, x_center float64, y_center float64) Solution {
    return Solution{radius, x_center, y_center}
}


func (this *Solution) RandPoint() []float64 {
    rand.Seed(time.Now().UnixNano())
    angle := 2 * math.Pi * rand.Float64() 
    sqrtRadius := math.Sqrt(rand.Float64()) * this.radius 
    x := this.x_center + sqrtRadius * math.Cos(angle)
    y := this.y_center + sqrtRadius * math.Sin(angle)
    return []float64{x, y}
}




/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(radius, x_center, y_center);
 * param_1 := obj.RandPoint();
 */
