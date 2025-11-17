func findMinimumOperations(s1 string, s2 string, s3 string) int {
    a, b, c := len(s1), len(s2), len(s3)
    n := min(a, b, c)
    
    if s1[0] != s2[0] || s1[0] != s3[0] {
        return -1
    }

    for i := 1; i < n; i++ {
        if s1[i] != s2[i] || s1[i] != s3[i] {
            return a + b + c - i * 3
        }
    }

    return a + b + c - n * 3
}