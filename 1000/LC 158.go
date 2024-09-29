/**
 * The read4 API is already defined for you.
 *
 *     read4 := func(buf4 []byte) int
 *
 * // Below is an example of how the read4 API can be called.
 * file := File("abcdefghijk") // File is "abcdefghijk", initially file pointer (fp) points to 'a'
 * buf4 := make([]byte, 4) // Create buffer with enough space to store characters
 * read4(buf4) // read4 returns 4. Now buf = ['a','b','c','d'], fp points to 'e'
 * read4(buf4) // read4 returns 4. Now buf = ['e','f','g','h'], fp points to 'i'
 * read4(buf4) // read4 returns 3. Now buf = ['i','j','k',...], fp points to end of file
 */

var solution = func(read4 func([]byte) int) func([]byte, int) int {
    var idx int = 0
    var rem int = 0
    var tmp []byte = make([]byte, 4)
    var f bool = true

    // implement read below.
    return func(buf []byte, n int) int {
        if f {
            rem = read4(tmp)
            f = false
        }

        i := 0
        for i < n && i < rem {
            buf[i] = tmp[idx % 4]
            i++
            idx++
            if idx % 4 == 0 {
                rem += read4(tmp)
            }
        }
        rem -= i
        return i 
    }
}