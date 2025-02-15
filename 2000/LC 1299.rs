impl Solution {
    pub fn replace_elements(mut arr: Vec<i32>) -> Vec<i32> {
        let n = arr.len();
        let mut suf = vec![-1; n];
        let mut mx = 0;
        for i in (1..n).rev() {
            mx = mx.max(arr[i]);
            suf[i-1] = mx;
        }

        suf
    }
}