func longestConsecutive(nums []int) int {
   
    numMap := make(map[int]bool)
    for _, v := range nums{
        numMap[v] = true
    }
    max := 0
    for _, v := range nums {
        if _, ok := numMap[v-1]; !ok {
            l := 1
            for numMap[v+l]{
                l++
            }
           if l > max {
            max = l
           }
        }
    } 
    return max


}
