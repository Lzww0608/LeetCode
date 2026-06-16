type BIT struct {
    n int 
    cnt, sum, sorted []int
}

func NewBIT(a []int) BIT {
    m := len(a)
    cnt := make([]int, m + 1)
    sum := make([]int, m + 1)
    n := 1 << (bits.Len(uint(m)) - 1)
    return BIT {
        n, cnt, sum, a,
    }
}

func (b *BIT)Update(i, num, val int) {
    for i < len(b.cnt) {
        b.cnt[i] += num 
        b.sum[i] += val
        i += i & (-i)
    }
}

func (b *BIT)GetKth(k int) int {
    i := 0
    for x := b.n; x > 0; x >>= 1 {
        nxt := i | x 
        if nxt < len(b.cnt) && b.cnt[nxt] < k {
            k -= b.cnt[nxt]
            i = nxt
        }
    }

    return b.sorted[i]
}

func (b *BIT)SumKth(k int) (res int) {
    i := 0
    for x := b.n; x > 0; x >>= 1 {
        nxt := i | x 
        if nxt < len(b.cnt) && b.cnt[nxt] < k {
            k -= b.cnt[nxt]
            res += b.sum[nxt]
            i = nxt
        }
    }

    return res + b.sorted[i] * k
}


func maxSum(nums []int, k int) int64 {
    n := len(nums)
    sorted := slices.Clone(nums)
    sort.Ints(sorted)
    sorted = slices.Compact(sorted)

    id := make([]int, n)
    allTree := NewBIT(sorted)
    sum := 0
    for i, x := range nums {
        id[i] = sort.SearchInts(sorted, x) + 1;
        allTree.Update(id[i], 1, x)
        sum += x
    }

    ans := math.MinInt
    inTree := NewBIT(sorted)
    outTree := NewBIT(sorted)
    for l := range n {
        clear(inTree.cnt)
        clear(inTree.sum)
        copy(outTree.cnt, allTree.cnt)
        copy(outTree.sum, allTree.sum)
        needSwap := 0
        cur := 0

        for r := l; r < n; r++ {
            x := nums[r]
            rk := id[r]
            cur += x 
            inTree.Update(rk, 1, x)
            outTree.Update(rk, -1, -x)

            inc := false 
            sz := r - l + 1 
            if needSwap < k && needSwap < sz && needSwap < n - sz {
                if inTree.GetKth(needSwap + 1) < outTree.GetKth(n - sz - needSwap) {
                    inc = true 
                    needSwap++
                }
            }

            if !inc && needSwap > 0 {
                if inTree.GetKth(needSwap) >= outTree.GetKth(n - sz - needSwap + 1) {
                    needSwap--
                }
            }

            delta := 0
            if needSwap > 0 {
                inSum := inTree.SumKth(needSwap)
                outSum := sum - cur - outTree.SumKth(n - sz - needSwap)
                delta = outSum - inSum
            }
            ans = max(ans, cur + delta)
        }
    }
    
    return int64(ans)
}