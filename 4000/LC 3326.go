const N int = 1_000_001

var primes [N]int

func init() {
    for i := 2; i < N; i++ {
        if primes[i] != 0 {
            continue
        }
        for j := i * i; j < N; j += i {
            if primes[j] == 0 {
                primes[j] = i
            }
        }
    }
}

func minOperations(nums []int) (ans int) {
    n := len(nums)

    for i := n - 2; i >= 0; i-- {
        if nums[i] <= nums[i + 1] {
            continue
        } else if primes[nums[i]] == 0 || primes[nums[i]] > nums[i + 1] {
            return -1
        }

        ans++
        nums[i] = primes[nums[i]]
    }

    return
}