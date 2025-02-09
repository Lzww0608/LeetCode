impl Solution {
    pub fn minimum_health(damage: Vec<i32>, armor: i32) -> i64 {
        let mut sum: i64 = 0;
        let mut mx: i32 = 0;
        for &x in &damage {
            sum += x as i64;
            mx = mx.max(x);
        }

        sum - mx.min(armor) as i64 + 1
    }
}