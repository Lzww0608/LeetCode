func subStrHash(s string, power int, modulo int, k int, hashValue int) string {
    n := len(s)
    hash, pk := 0, 1
    for i := n - 1; i >= n - k; i-- {
        hash = (hash * power + int(s[i] & 31)) % modulo
        pk = pk * power % modulo
    }

    ans := 0
    if hash == hashValue {
        ans = n - k
    }

    for i := n - k - 1; i >= 0; i-- {
        hash = (hash * power + int(s[i] & 31) - pk * int(s[i+k] & 31) % modulo + modulo) % modulo
        if hash == hashValue {
            ans = i
        }
    }

    return s[ans:ans+k]
}