use::std::cmp::min;
impl Solution {
    pub fn min_sum_of_lengths(arr: Vec<i32>, target: i32) -> i32 {
        let n = arr.len();
        let (mut l, mut sum) = (0, 0);
        let mut ans: i32 = std::i32::MAX;

        let mut f = vec![0; n + 1];
        f[0] = n + 1;

        for r in 0..n {
            sum += arr[r];

            while sum > target {
                sum -= arr[l];
                l += 1;
            }

            if sum == target {
                ans = min(ans, (r - l + 1 + f[l]) as i32);
                f[r+1] = min(f[r], r - l + 1);
            } else {
                f[r+1] = f[r];
            }
        } 

        if ans > n as i32 {
            return -1;
        }
        ans
    }
}