type Point struct {
	X float64
	Y float64
}

func intersection(start1, end1, start2, end2 []int) []float64 {

	if start1[0] > end1[0] {
		start1, end1 = end1, start1
	}
	if start2[0] > end2[0] {
		start2, end2 = end2, start2
	}

	if start1[0] > start2[0] {
		start1, end1, start2, end2 = start2, end2, start1, end1
	}

	A1 := end1[1] - start1[1]
	B1 := start1[0] - end1[0]
	C1 := end1[0]*start1[1] - start1[0]*end1[1]
	A2 := end2[1] - start2[1]
	B2 := start2[0] - end2[0]
	C2 := end2[0]*start2[1] - start2[0]*end2[1]
	determinant := A1*B2 - A2*B1
	determinantC := A1*C2 - A2*C1

	if determinant == 0 {
		if determinantC != 0 {
			return nil
		} else {
			if start1[0] != end1[0] {
				if end1[0] >= start2[0] {
					return []float64{float64(start2[0]), float64(start2[1])}
				}
				return nil
			}

			if start1[1] > end1[1] {
				start1, end1 = end1, start1
			}
			if start2[1] > end2[1] {
				start2, end2 = end2, start2
			}
			if start1[1] > start2[1] {
				start1, end1, start2, end2 = start2, end2, start1, end1
			}
			if end1[1] >= start2[1] {
				return []float64{float64(start2[0]), float64(start2[1])}
			}
			return nil
		}
	}

	x := float64(B1*C2-B2*C1) / float64(determinant)
	y := float64(A2*C1-A1*C2) / float64(determinant)

	if x >= float64(start1[0]) && x <= float64(end1[0]) && x >= float64(start2[0]) && x <= float64(end2[0]) &&
		y >= float64(min(start1[1], end1[1])) && y <= float64(max(start1[1], end1[1])) &&
		y >= float64(min(start2[1], end2[1])) && y <= float64(max(start2[1], end2[1])) {
		return []float64{x, y}
	}
	return nil
}

