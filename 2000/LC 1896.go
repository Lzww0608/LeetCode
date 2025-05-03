func minOperationsToFlip(s string) int {
    ops := []byte{}
    val := [][]int{}

    for i := range s {
        if s[i] == '0' || s[i] == '1' || s[i] == ')' {
            if s[i] == '0' {
                val = append(val, []int{0, 1})
            } else if s[i] == '1' {
                val = append(val, []int{1, 0})
            } else {
                ops = ops[:len(ops)-1]
            }

            if len(ops) > 0 && ops[len(ops)-1] != '(' {
                op := ops[len(ops)-1]
                ops = ops[:len(ops)-1]
                p1, q1 := val[len(val)-1][0], val[len(val)-1][1]
                val = val[:len(val)-1]
                p2, q2 := val[len(val)-1][0], val[len(val)-1][1]
                val = val[:len(val)-1]
                if op == '&' {
                    val = append(val, []int{min(p1, p2), min(q1 + q2, 1 + min(q1, q2))})
                } else {
                    val = append(val, []int{min(p1 + p2, 1 + min(p1, p2)), min(q1, q2)})
                }
            }
        } else {
            ops = append(ops, s[i])
        }
    }

    return max(val[len(val)-1][0], val[len(val)-1][1])
}