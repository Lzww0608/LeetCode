impl Solution {
    pub fn average_waiting_time(customers: Vec<Vec<i32>>) -> f64 {
        let mut ans: i64 = 0;
        let mut sum: i64 = 0;
        for v in customers.iter() {
            sum = sum.max(v[0] as i64) + v[1] as i64;
            ans += sum - v[0] as i64;
        }

        ans as f64 / customers.len() as f64
    }
}