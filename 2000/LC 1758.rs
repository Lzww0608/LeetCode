impl Solution {
    pub fn min_operations(s: String) -> i32 {
        let s = s.as_bytes();
        
        let check  = |x: u8| -> i32 {
            let mut cnt = 0;
            for i in 0..s.len() {
                let c = s[i] as u8;
                if (c - b'0') != x ^ (i & 1) as u8 {
                    cnt += 1;
                }
            }

            cnt
        };

        std::cmp::min(check(0), check(1))
    }
}