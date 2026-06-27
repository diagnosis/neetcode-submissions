func twoSum(numbers []int, target int) []int {
    l:= len(numbers)
    var left int 
    right := l - 1
    for left < right{
        if numbers[left] + numbers[right] == target{
            break
        }else if numbers[left] + numbers[right] > target {
            right--
        }else {
            left++
        }

    }
    return []int{left+1, right+1}

}
