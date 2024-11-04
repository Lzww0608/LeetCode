impl Solution {
    pub fn judge_square_sum(c: i32) -> bool {
        let mut a = 0i64;
        let mut b = (c as f64).sqrt() as i64;

        while a <= b {
            let s = a * a + b * b;
            if s == c as i64 {
                return true;
            } else if s < c as i64 {
                a += 1;
            } else {
                b -= 1;
            }
        }

        false
    }
}