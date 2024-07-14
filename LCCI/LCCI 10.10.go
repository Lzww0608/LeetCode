
type BIT struct {
    size int
    bit  []int
}


func NewBIT(size int) *BIT {
    return &BIT{
        size: size,
        bit:  make([]int, size+1),
    }
}


func (b *BIT) update(index, value int) {
    for index <= b.size {
        b.bit[index] += value
        index += index & -index
    }
}

func (b *BIT) query(index int) int {
    sum := 0
    for index > 0 {
        sum += b.bit[index]
        index -= index & -index
    }
    return sum
}

type StreamRank struct {
    bit *BIT
}

func Constructor() StreamRank {
    return StreamRank{
        bit: NewBIT(50001),
    }
}

func (sr *StreamRank) Track(x int) {
    sr.bit.update(x + 1, 1)
}

func (sr *StreamRank) GetRankOfNumber(x int) int {
    return sr.bit.query(x + 1)
}

