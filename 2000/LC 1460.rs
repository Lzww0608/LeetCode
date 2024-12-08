impl Solution {
    pub fn can_be_equal(target: Vec<i32>, arr: Vec<i32>) -> bool {
        let mut cnt = vec![0; 1001];
        for (&m, n) in target.iter().zip(arr) {
            cnt[m as usize] += 1;
            cnt[n as usize] -= 1;
        }

        cnt.iter().filter(|&x| *x != 0).count() == 0
    }
}