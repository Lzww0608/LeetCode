var f = [26]int{}

func find(x int) int {
    if x != f[x] {
        f[x] = find(f[x])
    }
    return f[x]
}

func merge(x, y int) {
    rx, ry := find(x), find(y)
    if rx < ry {
        f[ry] = rx
    } else if rx > ry {
        f[rx] = ry
    }
}

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
    for i := range f {
        f[i] = i
    }

    for i := range s1 {
        a, b := int(s1[i] - 'a'), int(s2[i] - 'a')
        merge(a, b)
    }

    builder := strings.Builder{}
    for _, c := range baseStr {
        x := int(c - 'a')
        builder.WriteByte(byte(find(x) + 'a'))
    }

    return builder.String()
}
