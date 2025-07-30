
import "reflect"
func closeStrings(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    cntS := make(map[int]int)
    cntT := make(map[int]int)
    for i := range s {
        x := int(s[i] - 'a')
        y := int(t[i] - 'a')
        cntS[x]++
        cntT[y]++
    }    
    for k := range cntS {
        if _, ok := cntT[k]; !ok {
            return false
        }
    }

    a := []int{}
    b := []int{}
    for _, v := range cntS {
        a = append(a, v)
    }
    for _, v := range cntT {
        b = append(b, v)
    }
    sort.Ints(a)
    sort.Ints(b)

    return reflect.DeepEqual(a, b)
}