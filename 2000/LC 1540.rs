impl Solution {
    pub fn can_convert_string(s: String, t: String, k: i32) -> bool {
        if s.len() != t.len() {
            return false;
        }

        let n = s.len();
        let mut cnt = vec![0; 26];
        let s = s.as_bytes();
        let t = t.as_bytes();
        for i in 0..n {
            if s[i] != t[i] {
                let x = (s[i] - b'a') as i32;
                let y = (t[i] - b'a') as i32;
                let d = (y - x + 26) % 26;
                if d + cnt[d as usize] * 26 > k {
                    return false;
                }
                cnt[d as usize] += 1;
            }
        }

        true
    }
}