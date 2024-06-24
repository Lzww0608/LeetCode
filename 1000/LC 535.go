type Codec struct {
    originToTiny map[string]string
    tinyToOrigin map[string]string
    str string
    pre string
    k int
}


func Constructor() Codec {
    return Codec {
        originToTiny: make(map[string]string),
        tinyToOrigin: make(map[string]string),
        str: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890",
        pre: "https://lzww.com/",
        k: 7,
    }
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	_, exists := this.originToTiny[longUrl]
    for !exists {
        cs := make([]byte, this.k)
        for i := 0; i < this.k; i++ {
            cs[i] = this.str[rand.Intn(len(this.str))];
        }
        cur := this.pre + longUrl
        if _, ok := this.tinyToOrigin[cur]; !ok {
            this.originToTiny[longUrl] = cur
            this.tinyToOrigin[cur] = longUrl
            break
        }
    }

    return this.originToTiny[longUrl]
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
    return this.tinyToOrigin[shortUrl]
}


/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
