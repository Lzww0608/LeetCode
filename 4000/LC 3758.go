func convertNumber(s string) string {
    m := map[string]byte {
        "zero": '0',
        "one": '1',
        "two": '2',
        "three": '3',
        "four": '4',
        "five": '5',
        "six": '6',
        "seven": '7',
        "eight": '8',
        "nine": '9',
    }
    
    ans := []byte{}
    n := len(s)
    for i := 0; i < n - 2; {
        if v, ok := m[s[i:i+3]]; ok {
            ans = append(ans, v)
            i += 3
            continue
        } 
        if i + 4 > n {
            i += 1
            continue
        }
        if v, ok := m[s[i:i+4]]; ok {
            ans = append(ans, v)
            i += 4
            continue
        } 
        if i + 5 > n {
            i += 1
            continue
        }
        if v, ok := m[s[i:i+5]]; ok {
            ans = append(ans, v)
            i += 5
            continue
        } 
        i += 1
    }

    return string(ans)
}