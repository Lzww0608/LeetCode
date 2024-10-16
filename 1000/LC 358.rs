use std::collections::BinaryHeap;
use std::collections::VecDeque;

impl Solution {
    pub fn rearrange_string(s: String, k: i32) -> String {
        if k > 26 {
            return String::new();
        }

        let n = s.len();
        let mut cnt = vec![0; 26];
        for &b in s.as_bytes() {
            let x = (b & 31) - 1;
            cnt[x as usize] += 1;
        }

        let mut ans = Vec::with_capacity(n);
        let mut heap = BinaryHeap::new();

        for (i, &x) in cnt.iter().enumerate() {
            if x > 0 {
                heap.push((x, i));
            }
        }

        let mut q = VecDeque::new();
        for _ in 0..n {
            if heap.is_empty() {
                return String::new();
            }

            let (mut t, id) = heap.pop().unwrap();
            ans.push((b'a' + id as u8) as char);
            t -= 1;
            q.push_back((t, id));

            if q.len() >= k as usize {
                let (t, id) = q.pop_front().unwrap();
                if t > 0 {
                    heap.push((t, id));
                }
            }
        }

        ans.into_iter().collect()
    }
}