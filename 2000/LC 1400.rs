impl Solution {
    pub fn can_construct(s: String, k: i32) -> bool {
        if s.len() < k as usize {
            return false;
        }

        let mut mask: i32 = 0;
        for c in s.chars() {
            let x: i32 = (c as u8 - b'a') as i32;
            mask ^= 1 << x;
        }

        let mut cnt: i32 = 0;
        for i in 0..32 {
            if mask & (1 << i) != 0 {
                cnt += 1;
            }
        }

        return cnt <= k;
    }
}