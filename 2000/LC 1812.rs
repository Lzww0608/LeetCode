impl Solution {
    pub fn square_is_white(coordinates: String) -> bool {
        let s = coordinates.as_bytes();
        let x = (s[0] - b'a') + (s[1] - b'1');

        x & 1 == 1
    }
}