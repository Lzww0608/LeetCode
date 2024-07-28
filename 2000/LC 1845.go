type SeatManager struct {
    sort.IntSlice // 继承 Len, Less, Swap
    seats int
}

func Constructor(int) SeatManager {
    return SeatManager{}
}

func (m *SeatManager) Reserve() int {
    if len(m.IntSlice) > 0 { 
        return heap.Pop(m).(int) 
    }
    m.seats += 1 
    return m.seats
}

func (m *SeatManager) Unreserve(seatNumber int) {
    heap.Push(m, seatNumber) 
}

func (m *SeatManager) Push(v any) { m.IntSlice = append(m.IntSlice, v.(int)) }
func (m *SeatManager) Pop() any   { a := m.IntSlice; v := a[len(a)-1]; m.IntSlice = a[:len(a)-1]; return v }