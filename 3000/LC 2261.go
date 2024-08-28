func countDistinct(nums []int, k int, p int) int {
    uniqueSubarrays := make(map[uint32]struct{})

	for i := range nums {
		count := 0
		subarray := []int{}
		for j := i; j < len(nums); j++ {
			subarray = append(subarray, nums[j])
			if nums[j] % p == 0 {
				count++
			}
			if count > k {
				break
			}
			subarrayHash := hashSubarray(subarray)
			uniqueSubarrays[subarrayHash] = struct{}{}
		}
	}

	return len(uniqueSubarrays)

}

func hashSubarray(subarray []int) uint32 {
	hasher := fnv.New32a()
	for _, num := range subarray {
		hasher.Write([]byte{byte(num)})
	}
	return hasher.Sum32()
}