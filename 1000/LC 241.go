
import "strconv"
func diffWaysToCompute(s string) []int {
    var merge func(s string) []int
    merge = func(s string) (res []int) {
        if len(s) <= 2 {
            t, _ := strconv.Atoi(s) 
            return []int{t}
        }

        for i := range s {
            if s[i] == '-' || s[i] == '+' || s[i] == '*' || s[i] == '%' {
                l := merge(s[:i])
                r := merge(s[i+1:])
                for _, x := range l {
                    for _, y := range r {
                        if s[i] == '-' {
                            res = append(res, x - y)
                        } else if s[i] == '+' {
                            res = append(res, x + y)
                        } else if s[i] == '*' {
                            res = append(res, x * y)
                        }
                    }
                }
            }
        }
        return
    }

    return merge(s)
}