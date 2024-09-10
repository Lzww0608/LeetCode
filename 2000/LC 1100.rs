impl Solution {
    pub fn num_k_len_substr_no_repeats(s: String, k: i32) -> i32 {
        let mut cnt = vec![0; 26];
        let mut ans = 0;
        let n = s.len();
        let (mut l, mut r) = (0, 0);

        while r < n {
            let x: u8 = s.chars().nth(r).unwrap() as u8 - b'a';
            cnt[x as usize] += 1;

            while cnt[x as usize] > 1 {
                let y: u8 = s.chars().nth(l).unwrap() as u8 - b'a';
                cnt[y as usize] -= 1;
                l += 1;
            }

            if r - l + 1 == k as usize {
                ans += 1;

                let y: u8 = s.chars().nth(l).unwrap() as u8 - b'a';
                cnt[y as usize] -= 1;
                l += 1;
            }

            r += 1;
        }

        ans
    }
}