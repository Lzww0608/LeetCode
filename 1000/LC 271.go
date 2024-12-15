
import "strconv"
type Codec struct {
}

// Encodes a list of strings to a single string.
func (c *Codec) Encode(strs []string) string {
    builder := strings.Builder{}
    for _, s := range strs {
        n := len(s)
        t := strconv.Itoa(n)
        for len(t) < 3 {
            t = "0" + t
        }
        builder.WriteString(t)
        builder.WriteString(s)
    }
    return builder.String()
}

// Decodes a single string to a list of strings.
func (c *Codec) Decode(strs string) (ans []string) {
    pos, n := 0, len(strs)
    for pos < n {
        t := strs[pos:pos+3]
        pos += 3
        x, _ := strconv.Atoi(t)
        ans = append(ans, strs[pos:pos+x])
        pos += x
    }
    return
}

// Your Codec object will be instantiated and called as such:
// var codec Codec
// codec.Decode(codec.Encode(strs));