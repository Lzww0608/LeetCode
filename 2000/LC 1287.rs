impl Solution {
    pub fn find_special_integer(arr: Vec<i32>) -> i32 {
        let n = arr.len();
        let mut i = 0;
        while i < n {
            let mut j = i + 1;
            while j < n && arr[j] == arr[i] {
                j += 1;
            }
            if j - i > n / 4 {
                return arr[i];
            }
            i = j;
        }

        -1
    }
}