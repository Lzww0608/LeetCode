impl Solution {
    pub fn missing_number(arr: Vec<i32>) -> i32 {
        let n = arr.len();
        let d = (arr[n-1] - arr[0]) / (n as i32);
        let (mut l, mut r) = (0, n - 1);
        while l + 1 < r {
            let mid = l + ((r - l) >> 1);
            if arr[mid] - arr[l] - d * (mid - l) as i32 != 0 {
                r = mid
            } else {
                l = mid
            }
        }

        (arr[l] + arr[r]) / 2
    }
}