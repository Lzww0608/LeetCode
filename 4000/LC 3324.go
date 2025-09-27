func stringSequence(target string) (ans []string) {
    pre := []byte{}
    for i := range target {
        pre = append(pre, 'a')
        ans = append(ans, string(pre))
        for j := 0; j < int(target[i] - 'a'); j++ {
            pre[len(pre) - 1] = byte('a' + j + 1)
            ans = append(ans, string(pre))
        }
    }

    return 
}