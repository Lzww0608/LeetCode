type pair struct {
    s string
    t int
}
type pair2 struct {
    x, y int
}
type UndergroundSystem struct {
    in map[int]pair
    out map[string]pair2
}


func Constructor() UndergroundSystem {
    return UndergroundSystem{map[int]pair{}, map[string]pair2{}}
}


func (this *UndergroundSystem) CheckIn(id int, stationName string, t int)  {
    this.in[id] = pair{stationName, t}
}


func (this *UndergroundSystem) CheckOut(id int, stationName string, t int)  {
    s := this.in[id].s + "to" + stationName
    if _, ok := this.out[s]; !ok {
        this.out[s] = pair2{0, 0}
    }
    this.out[s] = pair2{this.out[s].x + (t - this.in[id].t), this.out[s].y + 1}
}


func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
    s := startStation + "to" + endStation
    return float64(this.out[s].x) /float64(this.out[s].y) 
}


/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */
