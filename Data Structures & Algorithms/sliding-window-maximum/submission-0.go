func maxSlidingWindow(nums []int, k int) []int {
   deque := []int{}
   res := []int{}
   for i, v := range nums {
    for len(deque) > 0 && nums[deque[len(deque)-1]] < v {
        deque = deque[:len(deque)-1]
    }
    deque = append(deque, i)

     if deque[0] < i-k+1{
    deque = deque[1:]
   }

   if i >= k-1{
    res = append(res, nums[deque[0]])
   }
   }

  
   return res
}
