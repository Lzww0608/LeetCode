
type RangeModule struct {
	ranges []Interval
}

type Interval struct {
	start, end int
}

func Constructor() RangeModule {
	return RangeModule{ranges: []Interval{}}
}

func (this *RangeModule) AddRange(left int, right int) {
	newRanges := []Interval{}
	inserted := false

	for _, interval := range this.ranges {
		if right < interval.start {
			if !inserted {
				newRanges = append(newRanges, Interval{left, right})
				inserted = true
			}
			newRanges = append(newRanges, interval)
		} else if left > interval.end {
			newRanges = append(newRanges, interval)
		} else {
			left = min(left, interval.start)
			right = max(right, interval.end)
		}
	}

	if !inserted {
		newRanges = append(newRanges, Interval{left, right})
	}

	this.ranges = newRanges
}

func (this *RangeModule) QueryRange(left int, right int) bool {
	for _, interval := range this.ranges {
		if interval.start <= left && right <= interval.end {
			return true
		}
	}
	return false
}

func (this *RangeModule) RemoveRange(left int, right int) {
	newRanges := []Interval{}

	for _, interval := range this.ranges {
		if right <= interval.start || left >= interval.end {
			newRanges = append(newRanges, interval)
		} else {
			if interval.start < left {
				newRanges = append(newRanges, Interval{interval.start, left})
			}
			if interval.end > right {
				newRanges = append(newRanges, Interval{right, interval.end})
			}
		}
	}

	this.ranges = newRanges
}


