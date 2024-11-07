impl Solution {
    pub fn max_run_time(mut n: i32, mut batteries: Vec<i32>) -> i64 {
        let mut sum: i64 = 0;
        for &x in &batteries {
            sum += x as i64;
        }

        batteries.sort_by(|a, b| b.cmp(a));

        for &x in &batteries {
            if x as i64 > sum / n as i64 {
                sum -= x as i64;
                n -= 1;
            } else {
                return (sum / n as i64) as i64;
            }
        }

        0
    }
}