func maxArea(heights []int) int {
    var (
        maxArea int
        left int
        right int = len(heights) -1
    )

    for right > left{
        if heights[left] > heights[right]{
            if maxArea < heights[right] * (right - left){
                  maxArea = heights[right] * (right - left)
            }
          
            right--
        }else{
            if maxArea < heights[left] * (right - left){
                maxArea = heights[left] * (right - left)
            }
            left++
        }
        
        

    }
    return maxArea

}
