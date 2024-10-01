
impl Solution {
    pub fn min_sessions(tasks: Vec<i32>, session_time: i32) -> i32 {
        let n = tasks.len();
        let m = 1 << n;
        let mut sum = vec![0; m];
        for i in 0..n {
            let k = 1 << i;
            for j in 0..k {
                sum[j | k] = sum[j] + tasks[i];
            }
        }

        let mut f = vec![n; m];
        f[0] = 0;
        for s in 0..m {
            let mut sub = s;
            while sub > 0 {
                if sum[sub] <= session_time && f[s ^ sub] + 1 < f[s] {
                    f[s] = f[s ^ sub] + 1;
                }
                sub = (sub - 1) & s;
            }
        }

        f[m-1] as i32
    }
}