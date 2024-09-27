impl Solution {
    pub fn take_characters(s: String, k: i32) -> i32 {
        let mut cnt = [0; 3];
        for c in s.bytes() {
            cnt[(c - b'a') as usize] += 1;
        }

        if cnt[0] < k || cnt[1] < k || cnt[2] < k {
            return -1;
        }

        let mut mx = 0;
        let mut l = 0;
        for (r, &c) in s.as_bytes().iter().enumerate() {
            cnt[(c - b'a') as usize] -= 1;
            while cnt[(c - b'a') as usize] < k {
                cnt[(s.as_bytes()[l] - b'a') as usize] += 1;
                l += 1;
            }
            mx = mx.max(r - l + 1);
        }

        (s.len() - mx) as i32
    }
}
