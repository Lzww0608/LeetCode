impl Solution {
    pub fn closest_to_target(mut arr: Vec<i32>, k: i32) -> i32 {
        let mut ans: i32 = (arr[0] - k).abs();
        
        for i in 1..arr.len() {
            let x = arr[i];
            ans = std::cmp::min(ans, (x - k).abs());
            let mut j = i - 1;
            while j < arr.len() && arr[j] & x != arr[j] {
                arr[j] &= x;
                ans = ans.min((arr[j] - k).abs());
                j -= 1;
            }
        }

        ans
    }
}