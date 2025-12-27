package main

import (
	"math"
	"slices"
)

func defendSpaceCity(time []int, position []int) int {
	maxPos := slices.Max(position)
	meteors := make([]int, maxPos+1)
	for i, t := range time {
		p := position[i]
		meteors[p] |= 1 << (t - 1)
	}

	calc := func(mask, cost int) (res int) {
		if mask == 0 {
			return
		}

		for t := 0; t < 5; t++ {
			if mask&(1<<t) != 0 {
				d := 0
				for k := t; k < 5; k++ {
					if mask&(1<<k) != 0 {
						d++
						t = k
					} else {
						break
					}
				}
				res += cost + d
			}
		}

		return
	}

	const INF = math.MaxInt32
	f := make([]int, 32)
	for i := range f {
		f[i] = INF
	}
	f[0] = 0

	for _, x := range meteors {
		g := make([]int, len(f))
		for j := range g {
			g[j] = INF
		}

		for i := range 32 {
			if f[i] == INF {
				continue
			}
			for j := range 32 {
				if i&j != 0 {
					continue
				}

				k := i | j
				single := x & (^k)
				c1 := calc(j, 2)
				c2 := calc(single, 1)
				c := c1 + c2 + f[i]
				g[j] = min(g[j], c)
			}
		}

		f = g
	}

	return f[0]
}
