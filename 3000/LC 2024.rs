impl Solution {
    pub fn max_consecutive_answers(answer_key: String, k: i32) -> i32 {
        let n = answer_key.len();
        let bytes = answer_key.as_bytes();
        let mut ans = 0;

        let mut f = |c: u8| {
            let (mut l, mut r) = (0, 0);
            let mut cnt = 0;

            while r < n {
                if bytes[r] != c {
                    cnt += 1;
                }

                while cnt > k {
                    if bytes[l] != c {
                        cnt -= 1;
                    }
                    l += 1;
                }

                ans = ans.max((r - l + 1) as i32);
                r += 1;
            }
        };

        f(b'T');
        f(b'F');

        ans
    }
}