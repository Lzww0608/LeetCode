func minStickers(stickers []string, target string) int {
    n := len(target)
    f := make([]int, 1 << n)
    for i := range f {
        f[i] = math.MaxInt / 2
    }
    f[0] = 0

    for i := 0; i < len(f); i++ {
        if f[i] == math.MaxInt / 2 {
            continue
        }
        for _, s := range stickers {
            cur := i
            for _, c := range s {
                for j, t := range target {
                    if t == c && ((cur >> j) & 1) == 0 {
                        cur |= 1 << j
                        break
                    }
                }
            }
            f[cur] = min(f[cur], f[i] + 1)
        }
    }

    if f[len(f)-1] == math.MaxInt / 2 {
        return -1
    }
    return f[len(f)-1]
}