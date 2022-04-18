func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    
    for i, currentNumber := range nums {
        answer := target - currentNumber
        if j, found := m[answer]; found {
            return []int{i, j}
        }
        m[currentNumber] = i
    }
    return []int{}
}
