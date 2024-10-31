impl Solution {
    pub fn max_energy_boost(a: Vec<i32>, b: Vec<i32>) -> i64 {
        let n = a.len();
        let (mut maxA, mut maxB) = (a[0] as i64, b[0] as i64);

        for i in 1..n {
            let tmp = maxA;
            maxA = maxA.max(maxB - b[i-1] as i64) + a[i] as i64;
            maxB = maxB.max(tmp - a[i-1] as i64) + b[i] as i64;
        }

        maxA.max(maxB)
    }
}