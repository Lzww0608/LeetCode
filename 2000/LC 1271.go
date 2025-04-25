func toHexspeak(num string) string {
    a, _ := strconv.Atoi(num)
    b := []byte(fmt.Sprintf("%X", a))
    
    for i := range b {
        if b[i] >= '2' && b[i] <= '9' {
            return "ERROR"
        }
        if b[i] == '1' {
            b[i] = 'I'
        } else if b[i] == '0' {
            b[i] = 'O'
        }
    }
    
    return string(b)
}