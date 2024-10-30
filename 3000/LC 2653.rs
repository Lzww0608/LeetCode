impl Solution {
    pub fn get_subarray_beauty(nums: Vec<i32>, k: i32, y: i32) -> Vec<i32> {
        let n = nums.len();
        let mut ans: Vec<i32> = Vec::new();
        let mut cnt = vec![0; 101];

        for (i, &x) in nums.iter().enumerate() {
            cnt[x as usize + 50] += 1;
            if i as i32 >= k {
                cnt[nums[i - k as usize] as usize + 50] -= 1;
            }

            if i as i32 >= k - 1 {
                let mut cur = 0;
                let mut f = true;
                for i in 0..50 {
                    cur += cnt[i];
                    if cur >= y {
                        f = false;
                        ans.push(i as i32 - 50);
                        break;
                    }
                }
                if f {
                    ans.push(0);
                }
            }
        }

        ans
    }
}