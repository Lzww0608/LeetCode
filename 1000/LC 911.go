type TopVotedCandidate struct {
    times, res []int
}


func Constructor(persons []int, times []int) TopVotedCandidate {
    cnt := map[int]int{}
    n := len(persons)
    res := make([]int, n)
    mx, cur := 0, 0
    for i, x := range persons {
        cnt[x]++
        if cnt[x] >= mx {
            mx, cur = cnt[x], x
            res[i] = x
        } else {
            res[i] = cur
        }
    }

    return TopVotedCandidate{times, res}
}


func (this *TopVotedCandidate) Q(t int) int {
    l, r := 0, len(this.times) - 1
    for l < r {
        mid := r - ((r - l) >> 1)
        if this.times[mid] < t {
            l = mid
        } else if this.times[mid] > t {
            r = mid - 1
        } else {
            return this.res[mid]
        }
    }

    return this.res[l]
}


/**
 * Your TopVotedCandidate object will be instantiated and called as such:
 * obj := Constructor(persons, times);
 * param_1 := obj.Q(t);
 */