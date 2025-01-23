impl Solution {
    pub fn smallest_number(n: i32) -> i32 {
        let mut i = 1;
        while i < n {
            i = (i << 1) + 1;
        }
        i
    }
}