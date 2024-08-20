impl Solution {
    pub fn partition_string(s: String) -> i32 {
        let mut vis = 0;
        let mut ans: i32 = 1;
        for c in s.chars() {
            let x = c as u8 - b'a';
            if vis & (1 << x) != 0 {
                ans += 1;
                vis = 0;
            }
            vis |= 1 << x;
        }

        ans
    }
}