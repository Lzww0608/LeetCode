func mechanicalAccumulator(target int) (ans int) {
    var sum func(int) bool
    sum = func(target int) bool {
        ans += target
        return target > 0 && sum(target-1)
    }

    sum(target)
    return ans
}